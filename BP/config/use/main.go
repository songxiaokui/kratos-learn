package main

import (
	"encoding/json"
	"fmt"
	"kratos_learn/BP/config/parser"
	"kratos_learn/BP/config/parser/file"
	"kratos_learn/BP/config/pbconfig"
	"kratos_learn/BP/config/tools"
	"os"
)

func TestConvert(path string) {
	data, err := tools.Yaml2Json(path)
	if err != nil {
		fmt.Println("yaml2json error: ", err)
		return
	}

	var bc bpconfig.Bootstrap
	err = tools.Json2Proto(data, &bc)
	if err != nil {
		fmt.Println("json2proto error: ", err)
		return
	}

	d, _ := json.Marshal(&bc)
	fmt.Println("bc.name: ", bc.Name)
	fmt.Println("redis.db: ", bc.CacheRedis)
	_ = os.WriteFile("config.json", d, 0o666)
}

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
	fmt.Println("redis.db: ", bc.CacheRedis)
	TestConvert(configPath)

	select {}

}
