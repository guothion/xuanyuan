package common

import (
	"context"
	"strings"
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

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} { return c.ctx.Done() }

func (c *Context) Err() error { return c.ctx.Err() }

func (c *Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func (c *Context) WithServiceProfile(serviceProfile string) *Context {
	c.serviceProfile = strings.ToUpper(strings.TrimSpace(serviceProfile))
	return c
}

func (c *Context) WithAccessToken(token string) *Context {
	c.accessToken = token
	return c
}

func (c *Context) WithAccountName(accountName string) *Context {
	c.accountName = accountName
	return c
}

func (c *Context) ServiceProfile() string {
	return c.serviceProfile
}

func (c *Context) IsProd() bool {
	return c.serviceProfile == "PROD"
}

func (c *Context) IsDev() bool {
	return c.serviceProfile == "DEV"
}

func (c *Context) IsInte() bool {
	return c.serviceProfile == "INTE"
}

func (c *Context) IsQA() bool {
	return c.serviceProfile == "QA"
}

func (c *Context) AccessToken() string {
	return c.accessToken
}

func (c *Context) AccountName() string {
	return c.accountName
}

func (c *Context) Cancel() {
	c.cancelFunc()
}

func (c *Context) Inherit() *Context {
	return NewContext(time.Minute * 2).
		WithAccountName(c.accountName).
		WithAccessToken(c.accessToken).
		WithServiceProfile(c.serviceProfile)
}
