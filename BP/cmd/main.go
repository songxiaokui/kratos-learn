package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
	"test/server"
)

func newApp(gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID("test1"),
		kratos.Name("test2"),
		kratos.Version("v0.0.0"),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(log.NewStdLogger(os.Stdout)),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	hello := server.NewHelloServer()
	app := newApp(server.NewRPCService(hello), server.NewHTTPServer(hello))
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
