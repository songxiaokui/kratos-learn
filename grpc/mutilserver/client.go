package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"newgrpc/api"
)

const (
	address = "0.0.0.0:8006"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewHelloClient(conn)

	name := "ssss"
	md := make(metadata.MD, 0)
	md["content-type"] = []string{"application/grpc"}

	r, err := c.SayHello(context.Background(), &api.Req{Name: name}, grpc.Header(&md))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetData())
}
