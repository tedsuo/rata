package router

import (
	"fmt"
	"github.com/bmizerany/pat"
	"net/http"
	"net/url"
	"strings"
)

type Handlers map[string]http.Handler

type Params map[string]string

type Route struct {
	Handler string
	Method  string
	Path    string
}

func (r Route) PathWithParams(params Params) (string, error) {
	components := strings.Split(r.Path, "/")
	for i, c := range components {
		if len(c) == 0 {
			continue
		}
		if c[0] == ':' {
			val, ok := params[c[1:]]
			if !ok {
				return "", fmt.Errorf("missing param %s", c)
			}
			components[i] = val
		}
	}

	u, err := url.Parse(strings.Join(components, "/"))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

type Routes []Route

func (r Routes) RouteForHandler(handler string) (Route, bool) {
	for _, route := range r {
		if route.Handler == handler {
			return route, true
		}
	}
	return Route{}, false
}

func (r Routes) PathForHandler(handler string, params Params) (string, error) {
	route, ok := r.RouteForHandler(handler)
	if !ok {
		return "", fmt.Errorf("No route exists for handler %", handler)
	}
	return route.PathWithParams(params)
}

func (r Routes) Router(actions Handlers) (http.Handler, error) {
	p := pat.New()
	for _, route := range r {
		handler, ok := actions[route.Handler]
		if !ok {
			return nil, fmt.Errorf("missing handler %s", route.Handler)
		}
		switch strings.ToUpper(route.Method) {
		case "GET":
			p.Get(route.Path, handler)
		case "POST":
			p.Post(route.Path, handler)
		case "PUT":
			p.Put(route.Path, handler)
		case "DELETE":
			p.Del(route.Path, handler)
		default:
			return nil, fmt.Errorf("invalid verb: %s", route.Method)
		}
	}
	return p, nil
}
