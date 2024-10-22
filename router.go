package gorouter

import (
	"context"
	"net/http"
)

type Router struct {
	prefixPath  string
	routes      []RouteEntry
	middlewares []Middleware
}

type Middleware func(handler Handler) Handler

func NewRouter() *Router {
	return &Router{}
}

func (rtr *Router) Prefix(prefix string) *Router {
	rtr.prefixPath = prefix
	return rtr
}

func (rtr *Router) GET(path string, handler Handler) {
	rtr.addRouteEntry(http.MethodGet, path, handler)
}

func (rtr *Router) POST(path string, handler Handler) {
	rtr.addRouteEntry(http.MethodPost, path, handler)
}

func (rtr *Router) PUT(path string, handler Handler) {
	rtr.addRouteEntry(http.MethodPut, path, handler)
}

func (rtr *Router) DELETE(path string, handler Handler) {
	rtr.addRouteEntry(http.MethodDelete, path, handler)
}

func (rtr *Router) HandleFunc(method, path string, handler Handler) {
	rtr.addRouteEntry(method, path, handler)
}

func (rtr *Router) Routes(path string, fn MethodSetter) {
	gr := &GroupRoutes{path: rtr.prefixPath + path}

	fn(gr)

	if gr.routes != nil {
		rtr.routes = append(rtr.routes, gr.routes...)
	}
}

func (rtr *Router) Use(m ...Middleware) {
	rtr.middlewares = append(rtr.middlewares, m...)
}

// Listen listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
func (rtr *Router) Listen(port string) error {
	return http.ListenAndServe(port, rtr)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		writer:  w,
		request: r,
	}
	for _, route := range rtr.routes {
		params := route.match(r)
		if params == nil {
			continue
		}

		c := context.WithValue(ctx.request.Context(), "params", params)
		ctx.request = ctx.request.WithContext(c)

		rtr.execute(route, ctx)
		return
	}

	http.NotFound(w, r)
}

func (rtr *Router) execute(route RouteEntry, ctx *Context) {
	h := applyMiddleware(route.HandlerFunc, rtr.middlewares...)
	h(ctx)
}

func (rtr *Router) addRouteEntry(method, path string, handler Handler) {
	path = rtr.prefixPath + path
	exactPath := generatePath(path)

	rtr.routes = append(rtr.routes, RouteEntry{
		Method:      method,
		Path:        exactPath,
		HandlerFunc: handler,
	})
}

func applyMiddleware(h Handler, middleware ...Middleware) Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}
