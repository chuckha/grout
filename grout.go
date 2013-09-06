package grout

import (
	"net/http"
	"regexp"
)

type Route struct {
	pattern *regexp.Regexp
	handler RouteHandler
}

// pattern: /blogs/(?P<name>:[a-z][a-z_-]+[a-z])>/(?P<othername>:[0-9]+)
// When a request like /blogs/chucks-blog/200 comes in
// I want to match the this url to /blogs/[a-z][a-z_-]+[a-z]/[0-9]+
func NewRoute(pattern string, handler RouteHandler) (*Route, error) {
	route, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &Route{
		pattern: route,
		handler: handler,
	}, nil
}

type RouteHandler func(http.ResponseWriter, *http.Request, map[string]string)

type RouteMux struct {
	routes []*Route
}

func NewRouteMux() *RouteMux {
	routes := make([]*Route, 0)
	return &RouteMux{
		routes: routes,
	}
}

func (routes *RouteMux) Route(pattern string, handler RouteHandler) {
	newRoute, err := NewRoute(pattern, handler)
	if err != nil {
		panic(err)
	}
	routes.routes = append(routes.routes, newRoute)
}

func (routes *RouteMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	for _, route := range routes.routes {
		matches := route.pattern.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			names := route.pattern.SubexpNames()
			for i, name := range names {
				if i == 0 {
					continue
				}
				data[name] = matches[i]
			}
			route.handler(w, r, data)
			return
		}
	}
	http.NotFound(w, r)
}
