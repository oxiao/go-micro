package main

import (
	"fmt"

	"github.com/asim/go-micro/client"
	"github.com/asim/go-micro/transport"

	hello "github.com/asim/go-micro/examples/v4/greeter/srv/proto/hello"

	"context"
)

func init() {
	client.DefaultClient.Init(
		client.Transport(
			transport.NewHTTPTransport(transport.Secure(true)),
		),
	)
}

func main() {
	cl := hello.NewSayService("go.micro.srv.greeter", client.DefaultClient)

	rsp, err := cl.Hello(context.TODO(), &hello.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
