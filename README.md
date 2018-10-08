# aladdin

- A goroutine library written in go for game server.
- A micro-service library for game server.
- A RPC library for game server.

## 思路

1. 微服务架构。
1. 使用consul做服务发现。
1. 自定义协议生成（支持optional和required字段），使用之前的经验和积累，支持C#，Lua，Java和Go。
1. RPC支持coroutine的pause，同步代码，异步回调？
1. 支持容灾，链路超时配置，支持报警。
1. 完善的logger系统，参考log4j。
1. 存储使用redis和mongodb。
1. 客户端使用C#和Lua。
1. 模块化架构，可插拔。
1. 支持常用游戏服务：物品，背包，任务，排行榜，商店，商城，拍卖，活动，场景，战场。
1. 是否某些场景考虑无状态和有状态。
1. 延迟在200毫秒之内。
