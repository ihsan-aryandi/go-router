package rhaprouter

import (
	"net/http"
	"regexp"
)

func replacePathParams(path string) string {
	search := regexp.MustCompile("{(\\w+)}")
	result := search.ReplaceAllFunc([]byte(path), func(s []byte) []byte {
		group := search.ReplaceAllString(string(s), `(?P<$1>\w+)`)
		return []byte(group)
	})

	return string(result)
}

func isMethodMatches(r *http.Request, re *RouteEntry) bool {
	return r.Method == re.Method
}

func generatePath(path string) *regexp.Regexp {
	newPath := replacePathParams(path)
	return regexp.MustCompile("^" + newPath + "$")
}