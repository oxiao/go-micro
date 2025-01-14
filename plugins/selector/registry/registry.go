// Package registry uses the go-micro registry for selection
package registry

import (
	"github.com/asim/go-micro/cmd"
	"github.com/asim/go-micro/selector"
)

func init() {
	cmd.DefaultSelectors["registry"] = NewSelector
}

// NewSelector returns a new registry selector
func NewSelector(opts ...selector.Option) selector.Selector {
	return selector.NewSelector(opts...)
}
