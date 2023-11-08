package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "test/api"
)

func NewHTTPServer(hello *HelloServer) *http.Server {
	opts := make([]http.ServerOption, 0)
	opts = append(opts, http.Address("0.0.0.0:8003"))
	srv := http.NewServer(opts...)
	v1.RegisterHelloHTTPServer(srv, hello)
	return srv
}
