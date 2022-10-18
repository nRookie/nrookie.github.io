## 1 缓存雪崩、缓存击穿与缓存穿透





**缓存雪崩**：缓存在同一时刻全部失效，造成瞬时DB请求量大、压力骤增，引起雪崩。缓存雪崩通常因为缓存服务器宕机、缓存的 key 设置了相同的过期时间等引起。





**缓存击穿**：一个存在的key，在缓存过期的一刻，同时有大量的请求，这些请求都会击穿到 DB ，造成瞬时DB请求量大、压力骤增。







**缓存穿透**：查询一个不存在的数据，因为不存在则不会写到缓存中，所以每次都会去请求 DB，如果瞬间流量过大，穿透到 DB，导致宕机。



## Protobuf



https://geektutu.com/post/quick-go-protobuf.html

protobuf 的安装和使用教程请移步 [Go Protobuf 简明教程](https://geektutu.com/post/quick-go-protobuf.html)，这篇文章就不再赘述了。protobuf 广泛地应用于远程过程调用(RPC) 的二进制传输，使用 protobuf 的目的非常简单，为了获得更高的性能。传输前使用 protobuf 编码，接收方再进行解码，可以显著地降低二进制传输的大小。另外一方面，protobuf 可非常适合传输结构化数据，便于通信字段的扩展。

使用 protobuf 一般分为以下 2 步：

- 按照 protobuf 的语法，在 `.proto` 文件中定义数据结构，并使用 `protoc` 生成 Go 代码（`.proto` 文件是跨平台的，还可以生成 C、Java 等其他源码文件）。
- 在项目代码中引用生成的 Go 代码。







