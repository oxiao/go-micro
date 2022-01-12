package main

import (
	"github.com/asim/go-micro/cmd/micro/cli"

	// register commands
	_ "github.com/asim/go-micro/cmd/micro/cli/call"
	_ "github.com/asim/go-micro/cmd/micro/cli/describe"
	_ "github.com/asim/go-micro/cmd/micro/cli/generate"
	_ "github.com/asim/go-micro/cmd/micro/cli/new"
	_ "github.com/asim/go-micro/cmd/micro/cli/run"
	_ "github.com/asim/go-micro/cmd/micro/cli/services"
	_ "github.com/asim/go-micro/cmd/micro/cli/stream"
)

func main() {
	cli.Run()
}
