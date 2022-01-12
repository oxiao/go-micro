package util

import (
	"runtime/debug"

	"github.com/asim/go-micro/logger"
)

// GoSafe will run func in goroutine safely, avoid crash from unexpected panic
func GoSafe(fn func()) {
	if fn == nil {
		return
	}
	go func() {
		defer func() {
			if e := recover(); e != nil {
				logger.Errorf("[panic]%v\n%s", e, debug.Stack())
			}
		}()
		fn()
	}()
}
