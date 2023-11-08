package main

import "newgrpc/server"

func main() {
	go func() {
		server.RunRPCService()
	}()

	go func() {
		server.RunHttpService()
	}()

	select {}
}
