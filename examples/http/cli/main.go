package main

import (
	"context"
	"log"

	httpClient "github.com/asim/go-micro/plugins/client/http/v4"
	"github.com/asim/go-micro/client"
	"github.com/asim/go-micro/registry"
	"github.com/asim/go-micro/selector"
)

func main() {
	CallHttpServer()
}

func CallHttpServer() {
	r := registry.NewRegistry()
	s := selector.NewSelector(selector.Registry(r))
	// new client
	c := httpClient.NewClient(client.Selector(s))
	// create request/response
	request := c.NewRequest("demo-http", "/demo", "", client.WithContentType("application/json"))

	response := new(map[string]interface{})
	// call service
	err := c.Call(context.Background(), request, response)
	log.Printf("err:%v response:%#v\n", err, response)

}
