use futures::future::join_all;
use near_indexer::near_primitives::views::ReceiptView;
use near_indexer::near_primitives::{
    types::{AccountId, TransactionOrReceiptId},
    views::{ActionView, ReceiptEnumView},
};
use std::collections::HashMap;
use tokio::sync::mpsc;

use crate::errors::Result;

#[derive(Clone, Debug)]
pub(crate) struct CandidateData {
    pub rollup_id: u32,
    pub transaction_or_receipt_id: TransactionOrReceiptId,
    pub payloads: Vec<Vec<u8>>,
}

pub(crate) struct BlockListener {
    stream: mpsc::Receiver<near_indexer::StreamerMessage>,
    receipt_sender: mpsc::Sender<CandidateData>,
    addresses_to_rollup_ids: HashMap<AccountId, u32>,
}

#[derive(Clone, Debug, PartialEq, Eq, serde::Serialize, serde::Deserialize)]
pub(crate) struct ReceiptViewWithRollupId {
    pub(crate) rollup_id: u32,
    pub(crate) receipt_view: ReceiptView,
}

impl BlockListener {
    pub(crate) fn new(
        stream: mpsc::Receiver<near_indexer::StreamerMessage>,
        receipt_sender: mpsc::Sender<CandidateData>,
        addresses_to_rollup_ids: HashMap<AccountId, u32>,
    ) -> Self {
        Self {
            stream,
            receipt_sender,
            addresses_to_rollup_ids,
        }
    }

    fn receipt_filer_map(receipt: ReceiptViewWithRollupId) -> Option<CandidateData> {
        let actions = if let ReceiptEnumView::Action { actions, .. } = receipt.receipt_view.receipt {
            actions
        } else {
            return None;
        };

        let payloads = actions
            .into_iter()
            .filter_map(Self::extract_args)
            .collect::<Vec<Vec<u8>>>();

        if payloads.is_empty() {
            return None;
        }

        Some(CandidateData {
            rollup_id: receipt.rollup_id,
            transaction_or_receipt_id: TransactionOrReceiptId::Receipt {
                receipt_id: receipt.receipt_view.receipt_id,
                receiver_id: receipt.receipt_view.receiver_id,
            },
            payloads,
        })
    }

    fn extract_args(action: ActionView) -> Option<Vec<u8>> {
        match action {
            ActionView::FunctionCall { method_name, args, .. } if method_name == "submit" => Some(args.into()),
            _ => None,
        }
    }

    async fn process_stream(self) -> Result<()> {
        let Self {
            mut stream,
            receipt_sender,
            addresses_to_rollup_ids,
        } = self;

        while let Some(streamer_message) = stream.recv().await {
            // TODO: check receipt_receiver is closed?
            let candidates_data: Vec<CandidateData> = streamer_message
                .shards
                .into_iter()
                .flat_map(|shard| shard.chunk)
                .flat_map(|chunk| {
                    chunk
                        .receipts
                        .into_iter()
                        .filter_map(|receipt| {
                            if let Some(rollup_id) = addresses_to_rollup_ids.get(&receipt.receiver_id) {
                                Some(ReceiptViewWithRollupId {
                                    rollup_id: *rollup_id,
                                    receipt_view: receipt,
                                })
                            } else {
                                None
                            }
                        })
                        .filter_map(Self::receipt_filer_map)
                })
                .collect();

            if candidates_data.is_empty() {
                continue;
            }

            let results = join_all(candidates_data.into_iter().map(|receipt| receipt_sender.send(receipt))).await;
            results.into_iter().collect::<Result<_, _>>()?;
        }

        Ok(())
    }

    pub(crate) async fn start(self) -> Result<()> {
        let sender = self.receipt_sender.clone();
        tokio::select! {
            result = self.process_stream() => result,
            _ = sender.closed() => {
                Ok(())
            },
        }
    }
}

#[cfg(test)]
mod tests {
    use std::collections::HashMap;
    use crate::block_listener::{BlockListener, CandidateData, ReceiptViewWithRollupId};
    use near_crypto::{KeyType, PublicKey};
    use near_indexer::near_primitives::hash::CryptoHash;
    use near_indexer::near_primitives::types::{AccountId, TransactionOrReceiptId};
    use near_indexer::near_primitives::views::{ActionView, ReceiptEnumView, ReceiptView};
    use near_indexer::StreamerMessage;
    use std::path::PathBuf;
    use std::str::FromStr;
    use std::time::Duration;
    use tokio::sync::mpsc;
    use tokio::sync::mpsc::error::TryRecvError;

    impl PartialEq for CandidateData {
        fn eq(&self, other: &Self) -> bool {
            let res = match (&self.transaction_or_receipt_id, &other.transaction_or_receipt_id) {
                (
                    TransactionOrReceiptId::Receipt {
                        receiver_id,
                        receipt_id,
                    },
                    TransactionOrReceiptId::Receipt {
                        receipt_id: other_receipt_id,
                        receiver_id: other_receiver_id,
                    },
                ) => receipt_id == other_receipt_id && receiver_id == other_receiver_id,
                (
                    TransactionOrReceiptId::Transaction {
                        transaction_hash,
                        sender_id,
                    },
                    TransactionOrReceiptId::Transaction {
                        transaction_hash: other_hash,
                        sender_id: other_id,
                    },
                ) => transaction_hash == other_hash && sender_id == other_id,
                _ => false,
            };

            self.payloads == other.payloads && res
        }
    }

    #[test]
    fn test_candidate_data_extraction() {
        let rollup_id = 0;
        let da_contract_id = AccountId::from_str("a.test").unwrap();

        let test_predecessor_id = AccountId::from_str("test_predecessor").unwrap();
        let valid_receipt = ReceiptView {
            predecessor_id: test_predecessor_id.clone(),
            receiver_id: da_contract_id.clone(),
            receipt_id: CryptoHash::hash_bytes(b"test_receipt_id"),
            receipt: ReceiptEnumView::Action {
                signer_id: AccountId::from_str("test_signer").unwrap(),
                signer_public_key: PublicKey::empty(KeyType::ED25519),
                gas_price: 100,
                output_data_receivers: vec![],
                input_data_ids: vec![CryptoHash::hash_bytes(b"test_input_data_id")],
                actions: vec![ActionView::FunctionCall {
                    method_name: "submit".into(),
                    args: vec![1, 2, 3].into(),
                    gas: 100,
                    deposit: 100,
                }],
            },
        };
        let valid_receipt = ReceiptViewWithRollupId {
            rollup_id,
            receipt_view: valid_receipt,
        };

        let invalid_action_receipt = ReceiptView {
            predecessor_id: test_predecessor_id.clone(),
            receiver_id: da_contract_id.clone(),
            receipt_id: CryptoHash::hash_bytes(b"test_receipt_id"),
            receipt: ReceiptEnumView::Data {
                data_id: CryptoHash::hash_bytes(b"test_data_id"),
                data: Some(vec![1, 2, 3]),
            },
        };

        let invalid_action_receipt = ReceiptViewWithRollupId {
            rollup_id,
            receipt_view: invalid_action_receipt,
        };

        // Test valid receipt
        assert_eq!(
            BlockListener::receipt_filer_map(valid_receipt.clone()),
            Some(CandidateData {
                rollup_id,
                transaction_or_receipt_id: TransactionOrReceiptId::Receipt {
                    receipt_id: valid_receipt.receipt_view.receipt_id.clone(),
                    receiver_id: valid_receipt.receipt_view.receiver_id.clone(),
                },
                payloads: vec![vec![1, 2, 3]],
            })
        );

        // Test invalid action receipt
        assert_eq!(BlockListener::receipt_filer_map(invalid_action_receipt), None);
    }

    #[test]
    fn test_multiple_submit_actions() {
        let rollup_id = 10;
        let da_contract_id = AccountId::from_str("a.test").unwrap();
        let common_action_receipt = ReceiptEnumView::Action {
            signer_id: AccountId::from_str("test_signer").unwrap(),
            signer_public_key: PublicKey::empty(KeyType::ED25519),
            gas_price: 100,
            output_data_receivers: vec![],
            input_data_ids: vec![CryptoHash::hash_bytes(b"test_input_data_id")],
            actions: vec![
                ActionView::FunctionCall {
                    method_name: "submit".into(),
                    args: vec![1, 2, 3].into(),
                    gas: 100,
                    deposit: 100,
                },
                ActionView::FunctionCall {
                    method_name: "submit".into(),
                    args: vec![4, 4, 4].into(),
                    gas: 100,
                    deposit: 100,
                },
                ActionView::FunctionCall {
                    method_name: "random".into(),
                    args: vec![1, 2, 3].into(),
                    gas: 100,
                    deposit: 100,
                },
                ActionView::DeleteAccount {
                    beneficiary_id: da_contract_id.clone(),
                },
            ],
        };

        let test_predecessor_id = AccountId::from_str("test_predecessor").unwrap();
        let valid_receipt = ReceiptView {
            predecessor_id: test_predecessor_id.clone(),
            receiver_id: da_contract_id.clone(),
            receipt_id: CryptoHash::hash_bytes(b"test_receipt_id"),
            receipt: common_action_receipt.clone(),
        };
        let valid_receipt = ReceiptViewWithRollupId {
            rollup_id,
            receipt_view: valid_receipt,
        };

        // Test valid receipt
        assert_eq!(
            BlockListener::receipt_filer_map(valid_receipt.clone()),
            Some(CandidateData {
                rollup_id,
                transaction_or_receipt_id: TransactionOrReceiptId::Receipt {
                    receipt_id: valid_receipt.receipt_view.receipt_id.clone(),
                    receiver_id: valid_receipt.receipt_view.receiver_id.clone(),
                },
                payloads: vec![vec![1, 2, 3], vec![4, 4, 4]],
            })
        );
    }

    struct StreamerMessages {
        pub empty: Vec<StreamerMessage>,
        pub candidates: Vec<StreamerMessage>,
    }

    struct StreamerMessagesLoader;
    impl StreamerMessagesLoader {
        fn load() -> StreamerMessages {
            let test_data_dir = concat!(env!("CARGO_MANIFEST_DIR"), "/test_data");
            let empty_data_path = [test_data_dir, "/empty"].concat();
            let candidates_data_path = [test_data_dir, "/candidates"].concat();

            StreamerMessages {
                empty: Self::load_messages(&empty_data_path),
                candidates: Self::load_messages(&candidates_data_path),
            }
        }

        fn load_messages(dir_path: &str) -> Vec<StreamerMessage> {
            let files = std::fs::read_dir(dir_path)
                .map(|entry| {
                    entry
                        .map(|dir_entry| {
                            let dir_entry = dir_entry.unwrap();
                            dir_entry.path()
                        })
                        .collect::<Vec<PathBuf>>()
                })
                .unwrap();

            files
                .into_iter()
                .map(|file| {
                    let file = std::fs::File::open(file).unwrap();
                    serde_json::from_reader(file).unwrap()
                })
                .collect()
        }
    }

    #[tokio::test]
    async fn test_empty_listener() {
        let rollup_id = 1;
        let da_contract_id: AccountId = "da.test.near".parse().unwrap();

        let streamer_messages = StreamerMessagesLoader::load();

        let (sender, receiver) = mpsc::channel(10);
        let (receipt_sender, mut receipt_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(receiver, receipt_sender, HashMap::from([(da_contract_id, rollup_id)]));
        let _ = tokio::spawn(listener.start());

        for el in streamer_messages.empty {
            sender.send(el).await.unwrap();
        }

        tokio::time::sleep(Duration::from_millis(10)).await;

        assert_eq!(receipt_receiver.try_recv(), Err(TryRecvError::Empty));
    }

    #[tokio::test]
    async fn test_candidates_listener() {
        let rollup_id = 1;
        let da_contract_id = "da.test.near".parse().unwrap();
        let streamer_messages = StreamerMessagesLoader::load();

        let (sender, receiver) = mpsc::channel(10);
        let (receipt_sender, mut receipt_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(receiver, receipt_sender, HashMap::from([(da_contract_id, rollup_id)]));
        let handle = tokio::spawn(listener.start());

        let expected = streamer_messages.candidates.len();
        for el in streamer_messages.candidates {
            sender.send(el).await.unwrap();
        }

        drop(sender);
        handle.await.unwrap().unwrap();

        let mut counter = 0;
        while let Some(_) = receipt_receiver.recv().await {
            counter += 1;
        }

        assert_eq!(expected, counter);
    }

    #[tokio::test]
    async fn test_shutdown() {
        let rollup_id=5;
        let da_contract_id = "da.test.near".parse().unwrap();
        let streamer_messages = StreamerMessagesLoader::load();

        let (sender, receiver) = mpsc::channel(10);
        let (receipt_sender, receipt_receiver) = mpsc::channel(10);

        let listener = BlockListener::new(receiver, receipt_sender, HashMap::from([(da_contract_id, rollup_id)]));
        let handle = tokio::spawn(listener.start());
        for (i, el) in streamer_messages.empty.into_iter().enumerate() {
            sender.send(el).await.unwrap();

            // some random number
            if i == 5 {
                break;
            }
        }

        drop(receipt_receiver);
        // Sender::closed is trigerred
        assert!(handle.await.unwrap().is_ok());
    }
}
