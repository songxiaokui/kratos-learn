### 使用github.com/soheilhy/cmux 实现一个端口代理HTTP Websocket RPC

### 好用的在线客户端测试RPC工具
* 安装: go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
* 注册: 在grpc服务后注册即可
    ```text
    // Create your protocol servers.
 
	grpcS := grpc.NewServer()
	api.RegisterHelloServer(grpcS, server.NewHelloService())
	reflection.Register(grpcS)
    ```
* 开启服务:
  ```shell
    grpcui -plaintext 127.0.0.1:8006
  ```