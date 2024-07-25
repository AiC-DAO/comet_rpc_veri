# comet_rpc_veri

验证Comet节点如何从`JSON RPC`请求中恢复交易事务：

- 首先，节点收到`JSON RPC 2.0`的`Requst`字符串
- `Json`解析`Requst`字符串为语言相关的`Requst`类型
- `Base64`解析`Requst`类型`Params`中的`tx_bytes`字段(字符串)为字节数组
- `ProtoBuffer`解析字节数组为SDK的`proto`文件中定义的`Tx`类型
- 代码继续使用该`Tx`做其他工作（解析不同的`msg`等）


反向从客户端的角度考虑，如果想要构造一个发往`Comet RPC`的交易，那就需要反向构造：

`msgs👉Tx👉Tx bytes👉Tx bytes Base64 String👉Requst Params👉Requst👉Requst Json String`
