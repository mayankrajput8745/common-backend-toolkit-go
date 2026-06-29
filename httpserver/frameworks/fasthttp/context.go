package fasthttp

import (
	"encoding/json"

	"github.com/mayankrajput8745/common-backend-toolkit-go/httpserver/contract"

	fh "github.com/valyala/fasthttp"
)

type fasthttpContext struct {
	ctx *fh.RequestCtx
}

func newContext(ctx *fh.RequestCtx) contract.Context {
	return &fasthttpContext{ctx: ctx}
}

func (c *fasthttpContext) Method() string {
	return string(c.ctx.Method())

func (c *fasthttpContext) Path() string {
	return string(c.ctx.Path())
}

func (c *fasthttpContext) Param(key string) string {
	val := c.ctx.UserValue(key)
	if val == nil {
		return ""
	}
	s, _ := val.(string)
	return s
}

func (c *fasthttpContext) Query(key string) string {
	return string(c.ctx.QueryArgs().Peek(key))
}

func (c *fasthttpContext) Body() []byte {
	return c.ctx.PostBody()
}

func (c *fasthttpContext) Header(key string) string {
	return string(c.ctx.Request.Header.Peek(key))
}

func (c *fasthttpContext) StatusCode() int {
	return c.ctx.Response.StatusCode()
}

func (c *fasthttpContext) Status(code int) contract.Context {
	c.ctx.SetStatusCode(code)
	return c
}

func (c *fasthttpContext) WriteJSON(statusCode int, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.ctx.SetStatusCode(statusCode)
	c.ctx.SetContentType("application/json")
	c.ctx.SetBody(data)
	return nil
}

func (c *fasthttpContext) String(s string) error {
	c.ctx.SetBodyString(s)
	return nil
}

func (c *fasthttpContext) SetHeader(key, value string) {
	c.ctx.Response.Header.Set(key, value)
}
