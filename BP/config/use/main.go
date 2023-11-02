package main

import (
	"fmt"
	"kratos_learn/BP/config/parser"
	"kratos_learn/BP/config/parser/file"
	"kratos_learn/BP/config/pbconfig"
)

func main() {
	configPath := "/Users/austsxk/Desktop/修心/kratos-learn/BP/config/userdefine/cae_dev.yaml"

	c := config.New(
		config.WithSource(
			file.NewSource(configPath),
		),
	)

	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc bpconfig.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	fmt.Println("bc.name: ", bc.Name)
	fmt.Println("redis.db: ", bc.CacheRedis[0].DB.Value)
	select {}

}
