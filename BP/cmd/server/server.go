package server

import (
	"context"
	"fmt"
	v1 "test/api"
)

var _ v1.HelloHTTPServer = (*HelloServer)(nil)

type HelloServer struct {
	v1.UnimplementedHelloServer
}

func NewHelloServer() *HelloServer {
	return &HelloServer{}
}

func (hs *HelloServer) SayHello(c context.Context, req *v1.HelloReq) (*v1.Response, error) {
	return &v1.Response{Data: fmt.Sprintf("hello %s", req.Name)}, nil
}
