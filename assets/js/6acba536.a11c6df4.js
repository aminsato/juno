"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[408],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>k});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=r.createContext({}),p=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},u=function(e){var t=p(e.components);return r.createElement(s.Provider,{value:t},e.children)},c="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,s=e.parentName,u=l(e,["components","mdxType","originalType","parentName"]),c=p(n),d=a,k=c["".concat(s,".").concat(d)]||c[d]||m[d]||o;return n?r.createElement(k,i(i({ref:t},u),{},{components:n})):r.createElement(k,i({ref:t},u))}));function k(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=d;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l[c]="string"==typeof e?e:a,i[1]=l;for(var p=2;p<o;p++)i[p]=n[p];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},4370:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>i,default:()=>m,frontMatter:()=>o,metadata:()=>l,toc:()=>p});var r=n(7462),a=(n(7294),n(3905));const o={slug:"/",sidebar_position:1,title:"Quick Start"},i=void 0,l={unversionedId:"intro",id:"version-0.8.0/intro",title:"Quick Start",description:"Juno is your fast and featureful Starknet client implementation.",source:"@site/versioned_docs/version-0.8.0/intro.md",sourceDirName:".",slug:"/",permalink:"/",draft:!1,tags:[],version:"0.8.0",sidebarPosition:1,frontMatter:{slug:"/",sidebar_position:1,title:"Quick Start"},sidebar:"tutorialSidebar",next:{title:"Example Configuration",permalink:"/config"}},s={},p=[],u={toc:p},c="wrapper";function m(e){let{components:t,...n}=e;return(0,a.kt)(c,(0,r.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("p",null,(0,a.kt)("em",{parentName:"p"},"Juno is your fast and featureful Starknet client implementation.")),(0,a.kt)("p",null,"Suitable for casual setups, production-grade indexers, and everything in between."),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"\ud83d\udcbe ",(0,a.kt)("strong",{parentName:"li"},"Tiny database size"),": ~121Gb on mainnet"),(0,a.kt)("li",{parentName:"ul"},"\u26a1 ",(0,a.kt)("strong",{parentName:"li"},"Blazing fast sync"),": constrained only by hardware and the sequencer"),(0,a.kt)("li",{parentName:"ul"},"\ud83d\udcaf ",(0,a.kt)("strong",{parentName:"li"},"100% ",(0,a.kt)("a",{parentName:"strong",href:"https://github.com/starkware-libs/starknet-specs/tree/master"},"JSON-RPC spec")," compliance"),": all things Starknet, in one place"),(0,a.kt)("li",{parentName:"ul"},"\ud83c\udfce\ufe0f ",(0,a.kt)("strong",{parentName:"li"},"Minimal RPC response latency"),": to keep your applications moving"),(0,a.kt)("li",{parentName:"ul"},"\ud83d\udd0e ",(0,a.kt)("strong",{parentName:"li"},"Low-level GRPC database API"),": for the most demanding workloads")),(0,a.kt)("h1",{id:"sync-starknet-in-two-commands"},"Sync Starknet in Two Commands"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"# Juno's database directory. Can be any directory on the machine.\nmkdir -p junodb\n\n# Juno's HTTP server listens on port 6060.\ndocker run -d --name juno -p 6060:6060 -v junodb:/var/lib/juno nethermind/juno:latest --db-path /var/lib/juno --http --http-host 0.0.0.0\n")),(0,a.kt)("p",null,"For a complete list of options and their explanations, see the ",(0,a.kt)("a",{parentName:"p",href:"config"},"Example Configuration")," or run:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"docker run nethermind/juno --help\n")),(0,a.kt)("h1",{id:"juno-is-compatible-with-the-following-starknet-api-versions"},"Juno is compatible with the following Starknet API versions:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("strong",{parentName:"li"},"v0.6.0")," (Endpoint: ",(0,a.kt)("inlineCode",{parentName:"li"},"/v0_6"),")"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("strong",{parentName:"li"},"v0.5.0")," (Endpoint: ",(0,a.kt)("inlineCode",{parentName:"li"},"/v0_5"),")")),(0,a.kt)("p",null,"To interact with a specific API version, you can specify the version endpoint in your RPC calls. For example:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},'curl -X POST http://localhost:6060/v0_6 -H "Content-Type: application/json" --data \'{"jsonrpc":"2.0","method":"juno_version","id":1}\'\n')),(0,a.kt)("h1",{id:"looking-for-a-starknet-rpc-provider"},"Looking for a Starknet RPC Provider?"),(0,a.kt)("p",null,"Access Nethermind's Starknet RPC service for free at ",(0,a.kt)("a",{parentName:"p",href:"https://data.voyager.online"},"data.voyager.online"),"."),(0,a.kt)("h1",{id:"questions-discussions-community"},"Questions, Discussions, Community"),(0,a.kt)("p",null,"Find active Juno team members and users in the following places."),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://github.com/NethermindEth/juno"},"GitHub")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://discord.gg/SZkKcmmChJ"},"Discord")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://t.me/+LHRF4H8iQ3c5MDY0"},"Telegram"))))}m.isMDXComponent=!0}}]);