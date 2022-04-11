# go-jsonrpc-gateway

基于 grpc-gateway 魔改的 jsonrpc proxy。

## 初衷
随着公司业务逻辑愈加复杂，restful api 不能很好的满足业务需求，凭空带来了许多的沟通成本。

作为后端一直想让前端也能享受到 rpc 的便利。而 grpc-web 的弊端很明显，无法继承传统 http1.1 协议的生态，浏览器的支持程度也不够，用作前后端交互非常不合适。

利用 jsonrpc 协议是个比较符合现状的妥协方法，后端可以提供 grpc 接口的同时也可以复用之前的生态(比如 OAuth 2.0 等)。对于前端的侵入性也不大。