package main

import (
	"fmt"
	"kratos_learn/wire/pkg"
)

func main() {
	config := pkg.NewConfig()
	config.Sche.Type = "local"
	server, _ := initService(config)
	fmt.Println(server == nil)
	server.GetData()
	server.RunSched()
}
