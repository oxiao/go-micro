package client

import (
	"testing"

	"github.com/asim/go-micro/client"
)

func TestGetClient(t *testing.T) {
	s := &service{client: client.DefaultClient}
	if s.getClient("grpc").String() != "grpc" {
		t.Fail()
	}
	if s.getClient("http").String() != "http" {
		t.Fail()
	}
	if s.getClient("mucp").String() != "mucp" {
		t.Fail()
	}
	if s.getClient("other").String() != client.DefaultClient.String() {
		t.Fail()
	}
}
