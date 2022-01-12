package main

import (
	"fmt"

	yaml "github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/asim/go-micro/config"
	"github.com/asim/go-micro/config/reader"
	"github.com/asim/go-micro/config/reader/json"
	"github.com/asim/go-micro/config/source/file"
)

func main() {
	// new yaml encoder
	enc := yaml.NewEncoder()

	// new config
	c, _ := config.NewConfig(
		config.WithReader(
			json.NewReader( // json reader for internal config merge
				reader.WithEncoder(enc),
			),
		),
	)

	// load the config from a file source
	if err := c.Load(file.NewSource(
		file.WithPath("./config.yaml"),
	)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("data", c.Map())

	// define our own host type
	type Host struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}

	var host Host

	// read a database host
	if err := c.Get("hosts", "database").Scan(&host); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(host.Address, host.Port)
}
