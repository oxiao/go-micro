// Package mdns provides a multicast dns registry
package mdns

import (
	"github.com/asim/go-micro/cmd"
	"github.com/asim/go-micro/registry"
)

func init() {
	cmd.DefaultRegistries["mdns"] = NewRegistry
}

// NewRegistry returns a new mdns registry
func NewRegistry(opts ...registry.Option) registry.Registry {
	return registry.NewRegistry(opts...)
}

