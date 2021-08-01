## 基础

[[中英字幕\] 完整的gRPC课程（The complete gRPC course [Golang, Java, Protobuf]）_哔哩哔哩_bilibili](https://www.bilibili.com/video/BV1Xv411t7h5?p=6&spm_id_from=pageDriver)

### 简介

- Google 开源的一个高性能的 RPC(Remote Procedure Call) 框架
- 优点：
  - 提供高效的进程间通信。gRPC 没有使用 XML 或者 JSON 这种文本格式，而是采用了基于 protocol buffers 的二进制协议；同时，gRPC 采用了 HTTP/2 做为通信协议，从而能够快速的处理进程间通信
  - 简单且良好的服务接口和模式。gRPC 为程序开发提供了一种契约优先的方式，必须首先定义服务接口，才能处理实现细节
  - 支持多语言。gRPC 是语言中立的，我们可以选择任意一种编程语言，都能够与 gRPC 客户端或者服务端进行交互。
  - 成熟并且已被广泛使用。通过在 Google 的大量实战测试，gRPC 已经发展成熟
- 和`rest`的对比

![](C:\Users\yuxue\AppData\Roaming\Typora\typora-user-images\image-20210725234343015.png)

## ProtoBuf

## 序列化&反序列化ProtoBuf消息

## gRPC使用案例

## 其他

- idea的gradle配置
- protobuff那个不同语言的package配置，
- protobuff还是要多service间要维持一致的
- gRPC的不同使用模式
- go的init()方法
- 是用随机数来测试，testify，vscode中用颜色来显示的代码覆盖率
- uuid
- t.Prallel()
- 超时机制，错误参数检查，请求取消的实现
- 一元gRPC，就是一个clinet call 一个server；服务器流式rpc，客户端流式rpc，双向流式rpc
- 阻塞存根啥啥啥的
- gRPC反射，Evans客户端
- gRPC拦截器，验证和用户授权
- 利用SSL/TLS保护grpc连接
- grpc中的负载均衡
- gRPC网关
- go context
- struct tag
- 区分服务端和客户端的，客户端保留存根（stub），服务端构建grpc服务器
  - 服务端监听本机端口，创建一个grpc服务器实例，然后将protobuf中的service注册到grpc服务器上去
  - 客户端dial grpc服务器，获得连接，
- 一个servcie里面可以定义多个rpc方法
-  gRPC users typically call these APIs on the client side and implement the corresponding API on the server side.
- embeded struct
- slice计算长度的时间复杂度
- 两段代码的区别
- grouine实现map的协程安全
- C++的数组和go的切片不同

idea cmd + enter

感觉现在的机器和人交互方式没有很好利用语音

[Quick start | Go | gRPC](https://grpc.io/docs/languages/go/quickstart/)

一开始学习的最好材料还是官方文档

apigroup 中直接使用 datastore的读函数

对bittable的操作以函数的形式封装在datastore中

把数据转化的操作以函数形式封装在transforms，filters等中

然后APIgroup中直接使用这些函数就可以了