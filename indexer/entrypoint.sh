#!/bin/sh

/indexer-app/indexer init --chain-id localnet
/indexer-app/indexer run "$@"