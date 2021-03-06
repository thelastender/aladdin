# aladdin

- 使用Go语言实现的全套游戏服务器解决方案。

## 思路

1. 微服务架构。
1. 使用consul做服务发现。
1. 自定义协议生成（支持optional和required字段）/ protobuf，支持C#，Lua，Java和Go。
1. 支持容灾，链路超时配置，支持报警(opentsdb)。
1. 完善的logger系统，参考log4j。
1. 缓存和榜单存储使用redis。
1. 持久化存储使用TiDB/Cassandra/MySQL
1. 策划配置使用TOML/MySQL
1. 客户端使用C#和Lua。
1. 模块化架构，可插拔。
1. 支持常用游戏服务：物品，背包，任务，排行榜，商店，商城，拍卖，活动，场景，战场。
1. 某些场景无状态，某些场景有状态（动态更新固定模块）。
1. 延迟在200毫秒之内。
1. 服务治理

## 规范

- 命名采用非复数形式的命名方式，例如物品使用item，而不是items。
- [https://github.com/golang/go/wiki/CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)

## 配置

* 服务配置文件使用INI作为基本配置文件格式。
	- [https://ini.unknwon.io/](https://ini.unknwon.io/)
	- 因为极简，所以选择INI。
	- setting目录。

* 游戏配置（包括策划配置）使用protobuf来定义配置元数据，通过protoc生成读取和写入Excel、TOML、二进制的代码。其中，TOML作为中间可读和配置格式。
	* 策划配置是Excel格式，会导出成二进制。
	* 服务端一些配置是TOML格式，会导出成二进制。
	* 配置有Go读取写入Excel、TOML、二进制接口，C#读取二进制接口，Lua读取二进制接口。
	* 配置支持热更新，使用统一的配置接口。
	* config目录。

## 日志和打点

* 日志提供标准rsyslog的输出和本地文件的写入。

* metric使用UDP打点，每台机器上部署metric agent来收集打点数据，最终使用opentsdb存储打点数据。

## 通信协议

* 网络通信支持http协议（[Gin](https://github.com/gin-gonic/gin)），TCP，UDP，RUDP。
	* 协议使用protobuf，客户端使用C#的实现方案。
	* 加密协议使用lz4，客户端使用C#的实现方案。

* 服务RPC使用 自定义/protobuf 定义接口

## 服务治理

* 每一个服务有控制端口，支持指定的health检查和治理（超时配置，熔断，流量调度）




