package main

import (
	"fmt"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"newgrpc/api"
	"newgrpc/server"
)

type testHTTP1Handler struct {
}

type ExampleRPCRcvr struct {
}

func (h *testHTTP1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World")
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
func main() {
	// Create the main listener.
	l, err := net.Listen("tcp", ":8006")
	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux.
	m := cmux.New(l)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	//grpcL := m.Match(cmux.Any())
	httpL := m.Match(cmux.HTTP1Fast())
	trpcL := m.Match(cmux.Any()) // Any means anything that is not yet matched.

	// Create your protocol servers.
	grpcS := grpc.NewServer()
	api.RegisterHelloServer(grpcS, server.NewHelloService())
	reflection.Register(grpcS)

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	httpS := &http.Server{
		Handler: mux,
	}

	trpcS := rpc.NewServer()
	_ = trpcS.Register(&ExampleRPCRcvr{})

	// Use the muxed listeners for your servers.
	go func() {
		err := grpcS.Serve(grpcL)
		if err != nil {
			fmt.Println("RPC err: ", err)
		}
	}()

	go httpS.Serve(httpL)
	go trpcS.Accept(trpcL)
	// Start serving!
	m.Serve()

}
