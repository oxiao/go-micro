// Package mucp provides an mucp server
package mucp

import (
	"github.com/asim/go-micro/cmd"
	"github.com/asim/go-micro/server"
)

func init() {
	cmd.DefaultServers["mucp"] = NewServer
}

// NewServer returns a micro server interface
func NewServer(opts ...server.Option) server.Server {
	return server.NewServer(opts...)
}
