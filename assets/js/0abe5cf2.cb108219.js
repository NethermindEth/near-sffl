"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[849],{4270:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>d,contentTitle:()=>i,default:()=>p,frontMatter:()=>o,metadata:()=>a,toc:()=>l});var t=r(4848),s=r(8453);const o={sidebar_position:2},i="Setup",a={id:"operator/setup",title:"Setup",description:"While the opt-in process for NEAR SFFL is currently underway - see",source:"@site/docs/operator/setup.md",sourceDirName:"operator",slug:"/operator/setup",permalink:"/operator/setup",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:2,frontMatter:{sidebar_position:2},sidebar:"sidebar",previous:{title:"Registration",permalink:"/operator/registration"},next:{title:"Milestones",permalink:"/milestones"}},d={},l=[{value:"Hardware Requirements",id:"hardware-requirements",level:2},{value:"Minimal Hardware Specifications",id:"minimal-hardware-specifications",level:3},{value:"Steps",id:"steps",level:2},{value:"Step 1: Complete EigenLayer Operator Registration",id:"step-1-complete-eigenlayer-operator-registration",level:3},{value:"Step 2: Install Docker",id:"step-2-install-docker",level:3},{value:"Step 3: Prepare Local SFFL files",id:"step-3-prepare-local-sffl-files",level:3},{value:"Step 4: Copy your EigenLayer operator keys to the setup directory",id:"step-4-copy-your-eigenlayer-operator-keys-to-the-setup-directory",level:3},{value:"Step 5: Update your <code>.env</code> file",id:"step-5-update-your-env-file",level:3},{value:"Step 6: Update your configuration files",id:"step-6-update-your-configuration-files",level:3},{value:"Step 6: Set up your indexer",id:"step-6-set-up-your-indexer",level:3},{value:"Step 7: Synchronize your NEAR node",id:"step-7-synchronize-your-near-node",level:3},{value:"Step 7: Run your operator",id:"step-7-run-your-operator",level:3}];function c(e){const n={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",h3:"h3",p:"p",pre:"pre",strong:"strong",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,s.R)(),...e.components};return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)(n.h1,{id:"setup",children:"Setup"}),"\n",(0,t.jsx)(n.admonition,{type:"info",children:(0,t.jsxs)(n.p,{children:["While the opt-in process for NEAR SFFL is currently underway - see\n",(0,t.jsx)(n.a,{href:"./registration",children:"Registration"})," - the testnet is not completely operational\njust yet, so it's currently not required that operators run a node. Keep an\neye out for updates!"]})}),"\n",(0,t.jsx)(n.p,{children:"This guide will walk you through the steps required to set up your operator\nnode on the NEAR SFFL testnet. The testnet serves as a sandbox environment\nfor testing and development, allowing you to test both the AVS smart contracts\nand off-chain services. As the network is under active development, it's\ncrucial to stay updated with the latest changes and keep your node in sync\nwith the network."}),"\n",(0,t.jsx)(n.h2,{id:"hardware-requirements",children:"Hardware Requirements"}),"\n",(0,t.jsx)(n.p,{children:"A NEAR SFFL operator node consists of two main components: the AVS node\nsoftware and a NEAR DA indexer. The AVS node software is a Go implementation\nof the AVS protocol, while the NEAR DA indexer is essentially a NEAR full node\nthat indexes NEAR DA submissions on the NEAR blockchain."}),"\n",(0,t.jsx)(n.h3,{id:"minimal-hardware-specifications",children:"Minimal Hardware Specifications"}),"\n",(0,t.jsxs)(n.table,{children:[(0,t.jsx)(n.thead,{children:(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.th,{children:"Component"}),(0,t.jsx)(n.th,{children:"Specifications"})]})}),(0,t.jsxs)(n.tbody,{children:[(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"CPU"}),(0,t.jsx)(n.td,{children:"x86_64 (Intel, AMD) processor with at least 8 physical cores"})]}),(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"CPU Features"}),(0,t.jsx)(n.td,{children:"CMPXCHG16B, POPCNT, SSE4.1, SSE4.2, AVX"})]}),(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"RAM"}),(0,t.jsx)(n.td,{children:"32GB DDR4"})]}),(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"Storage"}),(0,t.jsx)(n.td,{children:"1.5TB SSD (NVMe SSD is recommended)"})]}),(0,t.jsxs)(n.tr,{children:[(0,t.jsx)(n.td,{children:"Operating System"}),(0,t.jsx)(n.td,{children:"Linux (tested on Ubuntu 22.04 LTS)"})]})]})]}),"\n",(0,t.jsx)(n.p,{children:"Verify CPU feature support by running the following command on Linux:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{children:'lscpu | grep -P \'(?=.*avx )(?=.*sse4.2 )(?=.*cx16 )(?=.*popcnt )\' > /dev/null \\\n  && echo "Supported" \\\n  || echo "Not supported"\n'})}),"\n",(0,t.jsx)(n.h2,{id:"steps",children:"Steps"}),"\n",(0,t.jsx)(n.admonition,{type:"note",children:(0,t.jsx)(n.p,{children:"At this initial testnet stage, operators need to be whitelisted. If you are\ninterested and have not already been whitelisted, please contact the SFFL\nteam!"})}),"\n",(0,t.jsx)(n.h3,{id:"step-1-complete-eigenlayer-operator-registration",children:"Step 1: Complete EigenLayer Operator Registration"}),"\n",(0,t.jsxs)(n.p,{children:["Complete the EigenLayer CLI installation and registration ",(0,t.jsx)(n.a,{href:"https://docs.eigenlayer.xyz/operator-guides/operator-installation",children:"here"}),"."]}),"\n",(0,t.jsx)(n.h3,{id:"step-2-install-docker",children:"Step 2: Install Docker"}),"\n",(0,t.jsxs)(n.p,{children:["Install ",(0,t.jsx)(n.a,{href:"https://docs.docker.com/engine/install/ubuntu/",children:"Docker Engine on Linux"}),"."]}),"\n",(0,t.jsx)(n.h3,{id:"step-3-prepare-local-sffl-files",children:"Step 3: Prepare Local SFFL files"}),"\n",(0,t.jsx)(n.p,{children:"Clone the SFFL repository and execute the following."}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{children:"git clone https://github.com/NethermindEth/near-sffl.git\ncd near-sffl/setup/operator\ncp .env.example .env\n"})}),"\n",(0,t.jsx)(n.h3,{id:"step-4-copy-your-eigenlayer-operator-keys-to-the-setup-directory",children:"Step 4: Copy your EigenLayer operator keys to the setup directory"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{children:"cp <path-to-your-operator-ecdsa-key> ./config/keys/ecdsa.json\ncp <path-to-your-operator-bls-key> ./config/keys/bls.json\n"})}),"\n",(0,t.jsxs)(n.h3,{id:"step-5-update-your-env-file",children:["Step 5: Update your ",(0,t.jsx)(n.code,{children:".env"})," file"]}),"\n",(0,t.jsxs)(n.p,{children:["You should have something similar to this in your ",(0,t.jsx)(n.code,{children:".env"}),":"]}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-bash",children:"# Tagged release for SFFL containers\nSFFL_RELEASE=latest\n\n# NEAR chain ID\nNEAR_CHAIN_ID=testnet\n\n# NEAR home and keys directories\nNEAR_HOME_DIR=~/.near\nNEAR_KEYS_DIR=~/.near-credentials\n\n# Operator BLS and ECDSA key passwords (from config/keys files)\nOPERATOR_BLS_KEY_PASSWORD=fDUMDLmBROwlzzPXyIcy\nOPERATOR_ECDSA_KEY_PASSWORD=EnJuncq01CiVk9UbuBYl\n"})}),"\n",(0,t.jsxs)(n.p,{children:["In general, you should set ",(0,t.jsx)(n.code,{children:"NEAR_HOME_DIR"})," and ",(0,t.jsx)(n.code,{children:"NEAR_KEYS_DIR"}),". Those are\nwhere your NEAR-related data will be stored. If you are using a block storage\nservice, you should set especially ",(0,t.jsx)(n.code,{children:"NEAR_HOME_DIR"})," to the block storage mount\npoint. Do note you should choose a directory that has enough space for your\nNEAR node's data, ",(0,t.jsx)(n.strong,{children:"which should be around 1TB"}),"."]}),"\n",(0,t.jsxs)(n.p,{children:["Then, set your EigenLayer ECDSA and BLS key passwords in the\n",(0,t.jsx)(n.code,{children:"OPERATOR_ECDSA_KEY_PASSWORD"})," and ",(0,t.jsx)(n.code,{children:"OPERATOR_BLS_KEY_PASSWORD"})," fields."]}),"\n",(0,t.jsx)(n.h3,{id:"step-6-update-your-configuration-files",children:"Step 6: Update your configuration files"}),"\n",(0,t.jsxs)(n.p,{children:["Now, in ",(0,t.jsx)(n.code,{children:"setup/operator/config/operator.yaml"}),", set the relevant fields."]}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-yaml",children:"# Production flag for logging - false for printing debug logs\nproduction: false\n\n# Operator ECDSA address\noperator_address: 0xD5A0359da7B310917d7760385516B2426E86ab7f\n\n# AVS contract addresses\navs_registry_coordinator_address: 0x0069A298e68c09B047E5447b3b762E42114a99a2\noperator_state_retriever_address: 0x8D0b27Df027bc5C41855Da352Da4B5B2C406c1F0\n\n# AVS network RPCs\n# *Important*: The WS RPC must allow event subscriptions. As Public Node\n# doesn't support it, you should use a different RPC provider.\neth_rpc_url: https://ethereum-holesky-rpc.publicnode.com\neth_ws_url: wss://ethereum-holesky-rpc.publicnode.com # You should change this!\n\n# EigenLayer ECDSA and BLS private key paths\necdsa_private_key_store_path: /near-sffl/config/keys/ecdsa.json\nbls_private_key_store_path: /near-sffl/config/keys/bls.json\n\n# Aggregator server IP and port\naggregator_server_ip_port_address: near-sffl-aggregator:8090\n\n# Operator EigenLayer metrics server IP and port\neigen_metrics_ip_port_address: 0.0.0.0:9090\nenable_metrics: true\nnode_api_ip_port_address: 0.0.0.0:9010\nenable_node_api: true\n\n# Whether to try and register the operator in the AVS and in EL on startup.\n# It will not re-register the operator if already registered.\n# If unset, the operator will not be registered on startup! You'll need to\n# manually register the operator - not recommended\nregister_operator_on_startup: true\n\n# RMQ address and indexer rollup IDs\nnear_da_indexer_rmq_ip_port_address: amqp://rmq:5672\nnear_da_indexer_rollup_ids: [421614, 11155420]\n\n# Rollup RPCs\nrollup_ids_to_rpc_urls:\n  421614: wss://arbitrum-sepolia-rpc.publicnode.com\n  11155420: wss://optimism-sepolia-rpc.publicnode.com\n\n# Token strategy address\n# Mock strategy to deposit when registering (only used for testing)\ntoken_strategy_addr: 0x0000000000000000000000000000000000000000\n"})}),"\n",(0,t.jsxs)(n.p,{children:["In general, you should first set your operator address in ",(0,t.jsx)(n.code,{children:"operator_address"}),",\nas well as your ",(0,t.jsx)(n.strong,{children:"Ethereum Holesky"})," RPC URLs in ",(0,t.jsx)(n.code,{children:"eth_rpc_url"})," and ",(0,t.jsx)(n.code,{children:"eth_ws_url"}),".\nPlease double-check that the WS RPC allows event subscriptions. We recommend\nthat you either use your own node's RPC or, in terms of providers, use Infura\nor Quicknode - Public Node unfortunately doesn't support it."]}),"\n",(0,t.jsxs)(n.p,{children:["Finally, set the aggregator server address in ",(0,t.jsx)(n.code,{children:"aggregator_server_ip_port_address"}),".\nYou should set this to the address that was sent to you during whitelisting."]}),"\n",(0,t.jsx)(n.p,{children:"It's also good to double-check all other configuration fields, such as the\ncontract addresses."}),"\n",(0,t.jsx)(n.h3,{id:"step-6-set-up-your-indexer",children:"Step 6: Set up your indexer"}),"\n",(0,t.jsx)(n.p,{children:"Follow the commands below in the operator setup directory:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{children:"source .env\ndocker compose --profile indexer up\n"})}),"\n",(0,t.jsx)(n.p,{children:"You should run it until it starts syncing. You'll see a log similar to:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{children:"near-sffl-indexer  | 2024-04-20T22:24:00.296255Z  INFO stats: #161784536 Waiting for peers 0 peers \u2b07 0 B/s \u2b06 0 B/s NaN bps 0 gas/s CPU: 0%, Mem: 13.7 GB\n"})}),"\n",(0,t.jsxs)(n.p,{children:["At this point, stop the execution with ",(0,t.jsx)(n.code,{children:"Ctrl+C"}),". We're going to use NEAR's\ndata snapshots to speed up the syncing process."]}),"\n",(0,t.jsx)(n.h3,{id:"step-7-synchronize-your-near-node",children:"Step 7: Synchronize your NEAR node"}),"\n",(0,t.jsxs)(n.p,{children:["In order to do that, follow the ",(0,t.jsx)(n.a,{href:"https://near-nodes.io/intro/node-data-snapshots",children:"instructions in NEAR Nodes"}),".\nDo remember that you'll need to download the snapshot to\n",(0,t.jsx)(n.code,{children:"${NEAR_HOME_DIR}/data"})," - based on your ",(0,t.jsx)(n.code,{children:".env"})," file."]}),"\n",(0,t.jsx)(n.p,{children:"After that, run the following again:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{children:"source .env\ndocker compose --profile indexer up\n"})}),"\n",(0,t.jsxs)(n.p,{children:["You most likely want to keep this as a separate screen, so you can use tools\nsuch as ",(0,t.jsx)(n.code,{children:"screen"})," or ",(0,t.jsx)(n.code,{children:"tmux"})," to keep it running as a separate session."]}),"\n",(0,t.jsx)(n.p,{children:"Your indexer should now continue the syncing process on it's own!\nKeep it running until it\u2019s time to run the operator, as it\u2019ll keep synced with\nNEAR."}),"\n",(0,t.jsx)(n.h3,{id:"step-7-run-your-operator",children:"Step 7: Run your operator"}),"\n",(0,t.jsx)(n.admonition,{type:"info",children:(0,t.jsx)(n.p,{children:"This step is only available once the testnet deployment is completely made."})}),"\n",(0,t.jsx)(n.p,{children:"This is the final step!"}),"\n",(0,t.jsxs)(n.p,{children:["Stop the previous execution with ",(0,t.jsx)(n.code,{children:"Ctrl+C"}),". Then, update your repository state:"]}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{children:"git stash\ngit pull\ngit stash pop\n"})}),"\n",(0,t.jsxs)(n.p,{children:["After that, double-check your ",(0,t.jsx)(n.code,{children:".env"})," and ",(0,t.jsx)(n.code,{children:"config/operator.yaml"})," files, then\nsimply run:"]}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{children:"source .env\ndocker compose --profile indexer --profile operator pull\ndocker compose --profile indexer --profile operator up\n"})}),"\n",(0,t.jsx)(n.p,{children:"Your operator node should now be up and running!"})]})}function p(e={}){const{wrapper:n}={...(0,s.R)(),...e.components};return n?(0,t.jsx)(n,{...e,children:(0,t.jsx)(c,{...e})}):c(e)}},8453:(e,n,r)=>{r.d(n,{R:()=>i,x:()=>a});var t=r(6540);const s={},o=t.createContext(s);function i(e){const n=t.useContext(o);return t.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function a(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:i(e.components),t.createElement(o.Provider,{value:n},e.children)}}}]);