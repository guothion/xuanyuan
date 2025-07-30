package common

import (
	"context"
	"time"
)

type Context struct {
	ctx            context.Context
	cancelFunc     context.CancelFunc
	serviceProfile string
	accessToken    string
	accountName    string
}

func NewContext(timeout time.Duration) *Context {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	return &Context{
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}
}
