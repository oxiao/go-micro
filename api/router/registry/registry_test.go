package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/asim/go-micro/registry"
)

func TestStoreRegex(t *testing.T) {
	router := newRouter()
	router.store([]*registry.Service{
		{
			Name:    "Foobar",
			Version: "latest",
			Endpoints: []*registry.Endpoint{
				{
					Name: "foo",
					Metadata: map[string]string{
						"endpoint":    "FooEndpoint",
						"description": "Some description",
						"method":      "POST",
						"path":        "^/foo/$",
						"handler":     "rpc",
					},
				},
			},
			Metadata: map[string]string{},
		},
	},
	)

	assert.Len(t, router.ceps["Foobar.foo"].pcreregs, 1)
}
