package gorouter

import (
	"context"
	"encoding/json"
	"net/http"
)

type Context struct {
	writer     http.ResponseWriter
	request    *http.Request
	statusCode int
}

type Map map[string]interface{}

// Request returns *http.Request
func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) Query(key string) string {
	return c.request.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	ctx := c.request.Context()
	params := ctx.Value("params").(map[string]string)
	return params[key]
}

func (c *Context) Body(v interface{}) error {
	return json.NewDecoder(c.request.Body).Decode(v)
}

func (c *Context) Cookies() []*http.Cookie {
	return c.request.Cookies()
}

func (c *Context) Cookie(name string) (*http.Cookie, error) {
	return c.request.Cookie(name)
}

func (c *Context) AddContext(key string, value interface{}) {
	ctx := context.WithValue(c.request.Context(), key, value)
	c.request = c.request.WithContext(ctx)
}

func (c *Context) GetContext(key string) interface{} {
	ctx := c.request.Context()
	return ctx.Value(key)
}

// ResponseWriter returns http.ResponseWriter
func (c *Context) ResponseWriter() http.ResponseWriter {
	return c.writer
}

func (c *Context) Header(key string) string {
	return c.writer.Header().Get(key)
}

func (c *Context) ContentType(value string) {
	c.writer.Header().Set("Content-Type", value)
}

func (c *Context) SetHeader(key, value string) {
	c.writer.Header().Set(key, value)
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.writer, cookie)
}

func (c *Context) Write(statusCode int, res string) {
	c.writer.Header().Set("Content-Type", "text/plain")
	c.writer.WriteHeader(statusCode)

	_, _ = c.writer.Write([]byte(res))
}

func (c *Context) JSON(statusCode int, v interface{}) {
	c.writer.Header().Set("Content-Type", "application/json")
	c.writer.WriteHeader(statusCode)

	_ = json.NewEncoder(c.writer).Encode(v)
}
