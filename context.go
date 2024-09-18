package rhaprouter

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Context struct {
	writer      http.ResponseWriter
	request     *http.Request
	statusCode  int
	requestTime time.Time
}

type Map map[string]interface{}

/*
	Request
*/
func (fwc *Context) Request() *http.Request {
	return fwc.request
}

func (fwc *Context) Query(key string) string {
	return fwc.request.URL.Query().Get(key)
}

func (fwc *Context) Param(key string) string {
	ctx := fwc.request.Context()
	params := ctx.Value("params").(map[string]string)
	return params[key]
}

func (fwc *Context) Body(v interface{}) error {
	return json.NewDecoder(fwc.request.Body).Decode(v)
}

func (fwc *Context) Cookies() []*http.Cookie {
	return fwc.request.Cookies()
}

func (fwc *Context) Cookie(name string) (*http.Cookie, error) {
	return fwc.request.Cookie(name)
}

func (fwc *Context) AddContext(key string, value interface{}) {
	c := context.WithValue(fwc.request.Context(), key, value)
	fwc.request = fwc.request.WithContext(c)
}

func (fwc *Context) GetContext(key string) interface{} {
	ctx := fwc.request.Context()
	return ctx.Value(key)
}

/*
	ResponseWriter
*/
func (fwc *Context) ResponseWriter() http.ResponseWriter {
	return fwc.writer
}

func (fwc *Context) Header(key string) string {
	return fwc.writer.Header().Get(key)
}

func (fwc *Context) ContentType(value string) {
	fwc.writer.Header().Set("Content-Type", value)
}

func (fwc *Context) StatusCode(statusCode int) *Context {
	fwc.statusCode = generateStatusCode(statusCode)
	return fwc
}

func (fwc *Context) SetHeader(key, value string) {
	fwc.writer.Header().Set(key, value)
}

func (fwc *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(fwc.writer, cookie)
}

func (fwc *Context) Write(res string) (int, error) {
	fwc.writer.Header().Set("Content-Type", "text/plain")
	fwc.writer.WriteHeader(fwc.getStatusCode())

	return fwc.writer.Write([]byte(res))
}

func (fwc *Context) JSON(v interface{}) error {
	fwc.writer.Header().Set("Content-Type", "application/json")
	fwc.writer.WriteHeader(fwc.getStatusCode())

	return json.NewEncoder(fwc.writer).Encode(v)
}

func (fwc *Context) getStatusCode() int {
	if fwc.statusCode == 0 {
		return http.StatusOK
	}
	return fwc.statusCode
}

func generateStatusCode(statusCode int) int {
	if statusCode == 0 {
		return http.StatusOK
	}
	return statusCode
}

/*
	Additional
*/
func (fwc *Context) RequestTime() time.Time {
	return fwc.requestTime
}