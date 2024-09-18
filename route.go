package rhaprouter

import (
	"net/http"
	"regexp"
)

type RouteEntry struct {
	Method      string
	Path        *regexp.Regexp
	HandlerFunc Handler
}

type Handler func(ctx *Context) error

func (re *RouteEntry) match(r *http.Request) map[string]string {
	if !isMethodMatches(r, re) {
		return nil
	}

	match := re.Path.FindStringSubmatch(r.URL.Path)
	if match == nil {
		return nil
	}

	params := make(map[string]string)
	groupNames := re.Path.SubexpNames()
	for i, group := range match {
		params[groupNames[i]] = group
	}

	return params
}
