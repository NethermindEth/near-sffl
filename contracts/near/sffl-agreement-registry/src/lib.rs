use std::str::FromStr;

use near_sdk::borsh::{self, BorshDeserialize, BorshSerialize};
use near_sdk::collections::{LookupMap, LookupSet, Vector};
use near_sdk::serde::{Deserialize, Serialize};
use near_sdk::{env, near_bindgen, AccountId};

use alloy_primitives::{Address, FixedBytes};
use alloy_sol_types::{eip712_domain, sol, Eip712Domain, SolStruct, SolType};

sol! {
    #[derive(Serialize, Deserialize)]
    #[serde(crate = "near_sdk::serde")]
    struct StateRootUpdateMessage {
        uint32 rollupId;
        uint64 blockHeight;
        uint64 nearBlockHeight;
        bytes32 stateRoot;
    }

    #[derive(Serialize, Deserialize)]
    #[serde(crate = "near_sdk::serde")]
    struct G1Point {
        uint256 X;
        uint256 Y;
    }

    #[derive(Serialize, Deserialize)]
    #[serde(crate = "near_sdk::serde")]
    struct Operator {
        G1Point pubkey;
        uint128 weight;
    }

    #[derive(Serialize, Deserialize)]
    #[serde(crate = "near_sdk::serde")]
    struct OperatorSetUpdateMessage {
        uint64 id;
        uint64 nearBlockHeight;
        Operator[] operators;
    }

    #[derive(Serialize, Deserialize)]
    #[serde(crate = "near_sdk::serde")]
    struct CheckpointTaskResponseMessage {
        uint32 referenceTaskIndex;
        bytes32 stateRootUpdatesRoot;
        bytes32 operatorSetUpdatesRoot;
    }

    #[derive(Serialize, Deserialize)]
    #[serde(crate = "near_sdk::serde")]
    struct EthNearAccountLink {
        address ethAddress;
        string nearAccountId;
    }
}

const DOMAIN: Eip712Domain = eip712_domain!(
    name: "SFFL",
    version: "0",
);

#[near_bindgen]
#[derive(BorshDeserialize, BorshSerialize)]
pub struct SFFLAgreementRegistry {
    operator_eth_address: LookupMap<AccountId, [u8; 20]>,
    state_root_updates: LookupMap<(u32, u64), Vector<Vec<u8>>>,
    operator_set_updates: LookupMap<u64, Vector<Vec<u8>>>,
    checkpoint_task_responses: LookupMap<u32, Vector<Vec<u8>>>,
    message_set: LookupSet<[u8; 32]>,
    message_signatures: LookupMap<[u8; 32], LookupMap<[u8; 20], [u8; 64]>>,
}

impl Default for SFFLAgreementRegistry {
    fn default() -> Self {
        Self {
            operator_eth_address: LookupMap::new(b"operator_eth_address".to_vec()),
            state_root_updates: LookupMap::new(b"state_root_updates".to_vec()),
            operator_set_updates: LookupMap::new(b"operator_set_updates".to_vec()),
            checkpoint_task_responses: LookupMap::new(b"checkpoint_task_responses".to_vec()),
            message_set: LookupSet::new(b"message_set".to_vec()),
            message_signatures: LookupMap::new(b"message_signatures".to_vec()),
        }
    }
}

#[near_bindgen]
impl SFFLAgreementRegistry {
    pub fn init_operator(&mut self, msg: &EthNearAccountLink, signature: &FixedBytes<64>, v: u8) {
        let signing_hash = msg.eip712_signing_hash(&DOMAIN);
        let address = self.ecrecover(&signing_hash, signature, v, true);

        if address != msg.ethAddress {
            std::panic!("Wrong message address");
        }

        let account_id = AccountId::from_str(&msg.nearAccountId)
            .unwrap_or_else(|_| std::panic!("Invalid account ID"));

        self.operator_eth_address
            .insert(&account_id, address.0.as_slice().try_into().unwrap());
    }

    pub fn post_state_root_update_signature(
        &mut self,
        msg: &StateRootUpdateMessage,
        signature: &FixedBytes<64>,
    ) {
        let eth_address = self.caller_eth_address();
        let encoded_msg = StateRootUpdateMessage::abi_encode_sequence(msg);
        let msg_hash = env::keccak256_array(&encoded_msg);

        if self.message_set.insert(&msg_hash) {
            self.state_root_updates
                .get(&(msg.rollupId, msg.blockHeight))
                .or_else(|| {
                    let prefix: Vec<u8> = [
                        b"state_root_updates_vec".as_slice(),
                        msg.rollupId.to_be_bytes().as_slice(),
                        msg.blockHeight.to_be_bytes().as_slice(),
                    ]
                    .concat();
                    let vec = Vector::new(prefix);

                    self.state_root_updates
                        .insert(&(msg.rollupId, msg.blockHeight), &vec);

                    Some(vec)
                })
                .unwrap()
                .push(&encoded_msg);
        }

        self.push_signature(&msg_hash, &eth_address, &signature.0);
    }

    pub fn post_operator_set_update_signature(
        &mut self,
        msg: &OperatorSetUpdateMessage,
        signature: &FixedBytes<64>,
    ) {
        let eth_address = self.caller_eth_address();
        let encoded_msg = OperatorSetUpdateMessage::abi_encode_sequence(msg);
        let msg_hash = env::keccak256_array(&encoded_msg);

        if self.message_set.insert(&msg_hash) {
            self.operator_set_updates
                .get(&msg.id)
                .or_else(|| {
                    let prefix: Vec<u8> = [
                        b"operator_set_updates_vec".as_slice(),
                        msg.id.to_be_bytes().as_slice(),
                    ]
                    .concat();
                    let vec = Vector::new(prefix);

                    self.operator_set_updates.insert(&msg.id, &vec);

                    Some(vec)
                })
                .unwrap()
                .push(&encoded_msg);
        }

        self.push_signature(&msg_hash, &eth_address, &signature.0);
    }

    pub fn post_checkpoint_signature(
        &mut self,
        msg: &CheckpointTaskResponseMessage,
        signature: &FixedBytes<64>,
    ) {
        let eth_address = self.caller_eth_address();
        let encoded_msg = CheckpointTaskResponseMessage::abi_encode_sequence(msg);
        let msg_hash = env::keccak256_array(&encoded_msg);

        if self.message_set.insert(&msg_hash) {
            self.checkpoint_task_responses
                .get(&msg.referenceTaskIndex)
                .or_else(|| {
                    let prefix: Vec<u8> = [
                        b"checkpoint_task_responses_vec".as_slice(),
                        msg.referenceTaskIndex.to_be_bytes().as_slice(),
                    ]
                    .concat();
                    let vec = Vector::new(prefix);

                    self.checkpoint_task_responses
                        .insert(&msg.referenceTaskIndex, &vec);

                    Some(vec)
                })
                .unwrap()
                .push(&encoded_msg);
        }

        self.push_signature(&msg_hash, &eth_address, &signature.0);
    }

    #[private]
    fn caller_eth_address(&self) -> [u8; 20] {
        self.operator_eth_address
            .get(&env::predecessor_account_id())
            .unwrap_or_else(|| std::panic!("No known ethereum address"))
    }

    #[private]
    fn push_signature(
        &mut self,
        msg_hash: &[u8; 32],
        eth_address: &[u8; 20],
        signature: &[u8; 64],
    ) {
        self.message_signatures
            .get(&msg_hash)
            .or_else(|| {
                let prefix: Vec<u8> =
                    [b"message_signatures_map".as_slice(), msg_hash.as_slice()].concat();

                let map = LookupMap::new(prefix);

                self.message_signatures.insert(&msg_hash, &map);

                Some(map)
            })
            .unwrap()
            .insert(&eth_address, signature.as_slice().try_into().unwrap());
    }

    #[private]
    fn ecrecover(
        &self,
        msg_hash: &FixedBytes<32>,
        signature: &FixedBytes<64>,
        v: u8,
        malleability_flag: bool,
    ) -> Address {
        match env::ecrecover(
            msg_hash.as_slice(),
            signature.as_slice(),
            v,
            malleability_flag,
        ) {
            Some(pubkey) => {
                let hash = env::keccak256_array(&pubkey);

                Address::from_slice(&hash[12..32])
            }
            None => std::panic!("Invalid signature"),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
}
