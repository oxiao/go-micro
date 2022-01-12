package main

import (
	"fmt"

	"github.com/asim/go-micro"
	"github.com/asim/go-micro/server"
)

func main() {
	service := micro.NewService()

	service.Server().Init(
		server.Wait(nil),
	)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
