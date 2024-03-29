use clap::Parser;
use configs::{Opts, SubCommand};
use tokio::sync::mpsc;
use tracing::error;

use crate::{
    block_listener::{BlockListener, CandidateData},
    candidates_validator::CandidatesValidator,
    configs::RunConfigArgs,
    errors::Result,
    rabbit_publisher::RabbitBuilder,
};

mod block_listener;
mod candidates_validator;
mod configs;
mod errors;
mod rabbit_publisher;

fn run(home_dir: std::path::PathBuf, config: RunConfigArgs) -> Result<()> {
    let addresses_to_rollup_ids = config.compile_addresses_to_ids_map()?;
    let rabbit_builder = RabbitBuilder::new(config.rmq_address);

    let indexer_config = near_indexer::IndexerConfig {
        home_dir,
        sync_mode: near_indexer::SyncModeEnum::FromInterruption,
        await_for_node_synced: near_indexer::AwaitForNodeSyncedEnum::WaitForFullSync,
        validate_genesis: true,
    };

    let system = actix::System::new();
    let block_res = system.block_on(async move {
        let indexer = near_indexer::Indexer::new(indexer_config).expect("Indexer::new()");
        let stream = indexer.streamer();
        let (view_client, _) = indexer.client_actors();

        // TODO: define buffer: usize const
        let (sender, receiver) = mpsc::channel::<CandidateData>(100);

        let rabbit_publisher = rabbit_builder.build()?;

        let block_listener = BlockListener::new(stream, sender, addresses_to_rollup_ids);
        let receipt_validator = CandidatesValidator::new(view_client, receiver, rabbit_publisher);

        let result = tokio::select! {
            result = receipt_validator.start() => result,
            result = block_listener.start() => result,
        };

        result
    });

    // Run until publishing finished
    system.run()?;

    block_res.map_err(|err| {
        error!(target: "sffl_indexer", "Indexer Error: {}", err);
        err
    })
}

fn main() -> Result<()> {
    // We use it to automatically search the for root certificates to perform HTTPS calls
    // (sending telemetry and downloading genesis)
    openssl_probe::init_ssl_cert_env_vars();
    let env_filter = near_o11y::tracing_subscriber::EnvFilter::new(
        "nearcore=info,publisher=info,tokio_reactor=info,near=info,\
         stats=info,telemetry=info,indexer=info,near-performance-metrics=info",
    );
    let _subscriber = near_o11y::default_subscriber(env_filter, &Default::default()).global();
    let opts: Opts = Opts::parse();

    let home_dir = opts.home_dir.unwrap_or(near_indexer::get_default_home());
    match opts.subcmd {
        SubCommand::Init(config) => near_indexer::indexer_init_configs(&home_dir, config.into())?,
        SubCommand::Run(config) => run(home_dir, config)?,
    }

    Ok(())
}
