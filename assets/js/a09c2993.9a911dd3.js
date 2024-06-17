"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[899],{1456:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>l,contentTitle:()=>r,default:()=>h,frontMatter:()=>a,metadata:()=>o,toc:()=>d});var s=t(4848),i=t(8453);const a={sidebar_position:1,slug:"/"},r="Nuffle Fast Finality Layer - NFFL",o={id:"introduction",title:"Nuffle Fast Finality Layer - NFFL",description:"Introduction",source:"@site/docs/introduction.md",sourceDirName:".",slug:"/",permalink:"/",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1,slug:"/"},sidebar:"sidebar",next:{title:"Protocol Design",permalink:"/category/protocol-design"}},l={},d=[{value:"Introduction",id:"introduction",level:2},{value:"Getting Started",id:"getting-started",level:2},{value:"Running step-by-step",id:"running-step-by-step",level:3},{value:"Dependencies",id:"dependencies",level:4},{value:"Steps",id:"steps",level:4},{value:"Running through Docker Compose",id:"running-through-docker-compose",level:3},{value:"Dependencies",id:"dependencies-1",level:4},{value:"Steps",id:"steps-1",level:4},{value:"More Details",id:"more-details",level:2}];function c(e){const n={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",h4:"h4",p:"p",pre:"pre",strong:"strong",...(0,i.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(n.h1,{id:"nuffle-fast-finality-layer---nffl",children:"Nuffle Fast Finality Layer - NFFL"}),"\n",(0,s.jsx)(n.h2,{id:"introduction",children:"Introduction"}),"\n",(0,s.jsx)(n.p,{children:"Rollups on the Ethereum network are gaining traction, indicating a new phase\nin the development of decentralized applications (dApps) and smart contracts.\nHowever, as the ecosystem continues to evolve towards a rollup-centric roadmap,\nit confronts new challenges such as state and liquidity fragmentation and\nextended finality time."}),"\n",(0,s.jsxs)(n.p,{children:["In order to solve this problem, the Nuffle Fast Finality Layer (NFFL, formerly SFFL) was\ndesigned. Through it, various chains can, while supplying block data to\n",(0,s.jsx)(n.a,{href:"https://github.com/near/rollup-data-availability",children:"NEAR DA"}),", rely on the\neconomic security of an ",(0,s.jsx)(n.a,{href:"https://www.eigenlayer.xyz",children:"EigenLayer"})," AVS to provide\na faster block finality to various protocols and use-cases while also including\nan additional public DA layer into their stack."]}),"\n",(0,s.jsx)(n.p,{children:"This universal, secure and fast finality leads to major advancements in\ninteroperability protocols, enabling or improving designs such as general\nbridging and chain abstraction."}),"\n",(0,s.jsxs)(n.p,{children:["For more details, refer to ",(0,s.jsx)(n.a,{href:"/design/overview",children:"Protocol Design"}),". NFFL is\nunder active development and is not yet available on any publicly\naccessible environments."]}),"\n",(0,s.jsx)(n.h2,{id:"getting-started",children:"Getting Started"}),"\n",(0,s.jsx)(n.h3,{id:"running-step-by-step",children:"Running step-by-step"}),"\n",(0,s.jsxs)(n.p,{children:["Through the project's ",(0,s.jsx)(n.code,{children:"make"})," scripts, you can set up each actor of the\nenvironment individually."]}),"\n",(0,s.jsx)(n.h4,{id:"dependencies",children:"Dependencies"}),"\n",(0,s.jsxs)(n.p,{children:["In order to set up the AVS environments, you'll first need to install\n",(0,s.jsx)(n.a,{href:"https://go.dev/dl/",children:"golang"}),",\n",(0,s.jsx)(n.a,{href:"https://doc.rust-lang.org/cargo/getting-started/installation.html",children:"rust"}),", and\n",(0,s.jsx)(n.a,{href:"https://nodejs.org/en/download",children:"node"}),".\nMake sure you're in a ",(0,s.jsx)(n.strong,{children:"unix environment"}),", as this is a pre-requisite\nfor running the NEAR indexer."]}),"\n",(0,s.jsxs)(n.p,{children:["Then, install ",(0,s.jsx)(n.a,{href:"https://book.getfoundry.sh/getting-started/installation",children:"foundry"}),",\n",(0,s.jsx)(n.code,{children:"go install"})," ",(0,s.jsx)(n.a,{href:"https://github.com/maoueh/zap-pretty",children:"zap-pretty"})," and ",(0,s.jsx)(n.code,{children:"npm install"}),"\n",(0,s.jsx)(n.a,{href:"https://github.com/near/near-cli",children:"near-cli v3"}),". One way of doing so would be:"]}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"curl -L https://foundry.paradigm.xyz | bash\nfoundryup\n\ngo install github.com/maoueh/zap-pretty@latest\nnpm install -g near-cli@3.5.0\n"})}),"\n",(0,s.jsxs)(n.p,{children:["You'll also need to install ",(0,s.jsx)(n.a,{href:"https://www.rabbitmq.com/docs/download",children:"RabbitMQ"}),"."]}),"\n",(0,s.jsx)(n.h4,{id:"steps",children:"Steps"}),"\n",(0,s.jsx)(n.p,{children:"First, initialize RabbitMQ. It will be necessary for the operator execution.\nThis can be a bit different depending on how it was installed."}),"\n",(0,s.jsx)(n.p,{children:"Then, start what should be the mainnet (i.e. AVS) network, with both EL and\nthe AVS contracts already deployed, and also the 'rollup' network:"}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"make start-anvil-chain-with-el-and-avs-deployed\n"})}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"make start-rollup-anvil-chain-with-avs-deployed\n"})}),"\n",(0,s.jsx)(n.p,{children:"Then, start the aggregator:"}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"make start-aggregator\n"})}),"\n",(0,s.jsx)(n.p,{children:"Then, start the indexer, which already executes a NEAR localnet, and set up\na NEAR DA contract:"}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"make start-indexer\n"})}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"make setup-near-da\n"})}),"\n",(0,s.jsx)(n.p,{children:"Lastly, start the operator and the relayer:"}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"make start-operator\n"})}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"make start-test-relayer\n"})}),"\n",(0,s.jsxs)(n.p,{children:["And that's it! You should be able to see each of the actors messaging each\nother as expected. You can edit some of the test parameters in the\n",(0,s.jsx)(n.code,{children:"/config-files"}),"."]}),"\n",(0,s.jsx)(n.h3,{id:"running-through-docker-compose",children:"Running through Docker Compose"}),"\n",(0,s.jsx)(n.p,{children:"You can also more easily run a similar testing environment through Docker\nCompose, in which each service is executed in a separate container."}),"\n",(0,s.jsx)(n.h4,{id:"dependencies-1",children:"Dependencies"}),"\n",(0,s.jsxs)(n.p,{children:["In order to build and run the containers, you'll need to install\n",(0,s.jsx)(n.a,{href:"https://www.docker.com/get-started/",children:"Docker"}),", as well as\n",(0,s.jsx)(n.a,{href:"https://ko.build/install/",children:"ko"}),"."]}),"\n",(0,s.jsxs)(n.p,{children:["You should also have ",(0,s.jsx)(n.code,{children:"make"})," for the build script, or examine and run the same\nsteps."]}),"\n",(0,s.jsx)(n.h4,{id:"steps-1",children:"Steps"}),"\n",(0,s.jsx)(n.p,{children:"First, build the containers:"}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"make docker-build-images\n"})}),"\n",(0,s.jsx)(n.p,{children:"Then, run:"}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"docker compose up\n"})}),"\n",(0,s.jsxs)(n.p,{children:["This will execute all services in the correct order and let you examine the\nindividual logs. You'll also be able to access each container's services from\nthe host through their image name, if necessary. The config files used for this\ntest are also at ",(0,s.jsx)(n.code,{children:"/config-files"}),", denominated with ",(0,s.jsx)(n.code,{children:"docker-compose"}),"."]}),"\n",(0,s.jsx)(n.p,{children:"To terminate all services, simply run:"}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-bash",children:"docker compose down\n"})}),"\n",(0,s.jsx)(n.h2,{id:"more-details",children:"More Details"}),"\n",(0,s.jsxs)(n.p,{children:["For more details, refer to ",(0,s.jsx)(n.a,{href:"/design/overview",children:"Protocol Design"}),"."]}),"\n",(0,s.jsxs)(n.p,{children:["The AVS implementation was based on the\n",(0,s.jsx)(n.a,{href:"https://github.com/Layr-Labs/incredible-squaring-avs",children:"Incredible Squaring AVS"}),"\nproject, from ",(0,s.jsx)(n.a,{href:"https://www.eigenlayer.xyz",children:"EigenLayer"}),"."]})]})}function h(e={}){const{wrapper:n}={...(0,i.R)(),...e.components};return n?(0,s.jsx)(n,{...e,children:(0,s.jsx)(c,{...e})}):c(e)}},8453:(e,n,t)=>{t.d(n,{R:()=>r,x:()=>o});var s=t(6540);const i={},a=s.createContext(i);function r(e){const n=s.useContext(a);return s.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function o(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:r(e.components),s.createElement(a.Provider,{value:n},e.children)}}}]);