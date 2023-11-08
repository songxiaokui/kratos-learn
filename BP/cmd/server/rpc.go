package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "test/api"
)

func NewRPCService(hello *HelloServer) *grpc.Server {
	opts := make([]grpc.ServerOption, 0)
	opts = append(opts, grpc.Address("0.0.0.0:8001"))
	hs := grpc.NewServer(opts...)
	v1.RegisterHelloServer(hs, hello)
	return hs
}
