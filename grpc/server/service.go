package server

import (
	"context"
	"fmt"
	"newgrpc/api"
)

type HelloService struct {
	api.UnimplementedHelloServer
}

func (hs *HelloService) SayHello(c context.Context, req *api.Req) (*api.Response, error) {
	return &api.Response{
		Data: fmt.Sprintf("你好呀 %s!", req.Name),
	}, nil
}

func NewHelloService() *HelloService {
	return &HelloService{}
}
