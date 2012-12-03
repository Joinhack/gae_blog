package myapp

import (
	"strings"
	"errors"
//	"fmt"

	"net/http"
)

var (
	NotMatch = errors.New("not match the parameters")
)

type HandleFunc func(w http.ResponseWriter, r *http.Request, pathParam map[string]string)

type Handler struct {
	Express string
	IsRest bool
	Paths []string
	Function HandleFunc
}

type Route struct {
	Handlers []*Handler
}

func (h *Handler) ExtractParams(path string, s []string) (param map[string]string, err error) {
	err = NotMatch
	param = make(map[string]string)
	if !h.IsRest {
		if path == h.Express {
			return param, nil
		} else {
			return param, err
		}
	}
	if len(s) != len(h.Paths) {
		return param, err
	}
	for i, path := range h.Paths {
		idx := 0
		var tmp string
		tmp = path[idx:len(path)]
		if idx = strings.Index(tmp, "{"); idx == -1 && tmp != s[i] {
			if len(tmp) == len(path)  {
				return param, err
			} else {
					continue
			}
		} else if idx == -1 && tmp == s[i] {
			continue
		}
		pidx := idx
		idx = strings.Index(tmp, "}")
		if idx == -1 {
			return param, err
		}
		name := tmp[pidx + 1:idx]
		var value string
		if idx + 1 == len(tmp) {
			value = s[i][pidx:len(s[i])]
		} else {
			end := tmp[idx + 1: len(tmp)]
			idx = strings.LastIndex(s[i], end)
			if idx == -1 {
				return param, err
			}
			value = s[i][pidx: idx]
			param[name] = value
		}
	}
	return param, nil
}

func init() {
	http.HandleFunc("/", Dispatch)
}

var routes []*Route = make([]*Route, 0, 10)

func NewRoute() *Route {
	r := new(Route)
	routes = append(routes, r)
	return r
}

func (r *Route) HandleFunc(express string, function HandleFunc) {
	h := new(Handler)
	h.Express = express
	h.Function = function
	if idx := strings.Index(express, "{"); idx == -1 {
		h.IsRest = false
	} else {
		h.IsRest = true
	}
	h.Paths = strings.Split(express, "/")
	r.Handlers = append(r.Handlers, h)
}

func Dispatch(w http.ResponseWriter, r *http.Request) {
	uriFull := r.URL.RequestURI()
	uri := strings.Split(uriFull, "?")[0]
	uriPaths := strings.Split(uri, "/")
	for _, route := range routes {
		for _, handle := range route.Handlers {
			params, err := handle.ExtractParams(uri, uriPaths)
			if err == nil {
				handle.Function(w, r, params)
				return
			}
		}
	}
}


