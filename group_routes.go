package gorouter

import (
	"net/http"
)

type GroupRoutes struct {
	path   string
	routes []RouteEntry
}

type MethodSetter func(route *GroupRoutes)

func (gr *GroupRoutes) On(method string, handler Handler) {
	gr.generateRouteEntry(method, handler)
}

func (gr *GroupRoutes) GET(handler Handler) {
	gr.generateRouteEntry(http.MethodGet, handler)
}

func (gr *GroupRoutes) POST(handler Handler) {
	gr.generateRouteEntry(http.MethodPost, handler)
}

func (gr *GroupRoutes) PUT(handler Handler) {
	gr.generateRouteEntry(http.MethodPut, handler)
}

func (gr *GroupRoutes) DELETE(handler Handler) {
	gr.generateRouteEntry(http.MethodDelete, handler)
}

func (gr *GroupRoutes) generateRouteEntry(method string, handler Handler) {
	exactPath := generatePath(gr.path)

	gr.routes = append(gr.routes, RouteEntry{
		Method:      method,
		Path:        exactPath,
		HandlerFunc: handler,
	})
}
