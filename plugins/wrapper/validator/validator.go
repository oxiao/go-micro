package validator

import (
	"context"

	"github.com/asim/go-micro/errors"
	"github.com/asim/go-micro/server"
)

type Validator interface {
	Validate() error
}

func NewHandlerWrapper() server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			if v, ok := req.Body().(Validator); ok {
				if err := v.Validate(); err != nil {
					return errors.BadRequest(req.Service(), "%v", err)
				}
			}
			return fn(ctx, req, rsp)
		}
	}
}
