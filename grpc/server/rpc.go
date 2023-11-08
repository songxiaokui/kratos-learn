package server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"newgrpc/api"
)

func RunRPCService() {
	listen, _ := net.Listen("tcp", "0.0.0.0:8004")
	server := grpc.NewServer()
	api.RegisterHelloServer(server, NewHelloService())
	fmt.Println("rpc server start successful, listen on address: ", "0.0.0.0:8004")
	_ = server.Serve(listen)
}
