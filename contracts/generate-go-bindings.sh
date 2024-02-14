#!/bin/bash

function create_binding {
    contract_dir=$1
    contract=$2
    binding_dir=$3
    echo "generating bindings for" $contract
    mkdir -p $binding_dir/${contract}
    contract_json="$contract_dir/evm/out/${contract}.sol/${contract}.json"
    solc_abi=$(cat ${contract_json} | jq -r '.abi')
    solc_bin=$(cat ${contract_json} | jq -r '.bytecode.object')

    mkdir -p tmp
    echo ${solc_abi} >tmp/tmp.abi
    echo ${solc_bin} >tmp/tmp.bin

    rm -f $binding_dir/${contract}/binding.go
    abigen --bin=tmp/tmp.bin --abi=tmp/tmp.abi --pkg=contract${contract} --out=$binding_dir/${contract}/binding.go
    rm -rf tmp
}

rm -rf bindings/*
forge clean --root ./evm
forge build --root ./evm

avs_service_contracts="SFFLServiceManager SFFLTaskManager SFFLRegistryRollup SFFLRegistryCoordinator SFFLOperatorSetUpdateRegistry"
for contract in $avs_service_contracts; do
    create_binding . $contract ./bindings
done

create_binding . ERC20Mock ./bindings
