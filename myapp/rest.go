package myapp

import (
	"strings"
	"errors"
	"fmt"

	"net/http"
)

var (
	NotMatch = errors.New("not match the parameters")
)

type HandleFunc func(w http.ResponseWriter, r *http.Request, pathParam map[string]string)

type Handler struct {
	Express string
	Paths []string
	Function HandleFunc
}

type Route struct {
	Handlers []*Handler
}

func (h *Handler) ExtractParams(s []string) (param map[string]string, err error) {
	err = NotMatch
	param = make(map[string]string)
	if len(s) != len(h.Paths) {
		return param, err
	}
	for i, path := range h.Paths {
		idx := 0
		var tmp string
		//label:
		tmp = path[idx:len(path)]
		fmt.Println("-----------", len(path), s[i])
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
		println(name)
	}
	return
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
	h.Paths = strings.Split(express, "/")
	r.Handlers = append(r.Handlers, h)
}

func Dispatch(w http.ResponseWriter, r *http.Request) {
	uriFull := r.URL.RequestURI()
	uri := strings.Split(uriFull, "?")[0]
	uriPaths := strings.Split(uri, "/")
	println(uri)
	for _, route := range routes {
		for _, handle := range route.Handlers {
			params, err := handle.ExtractParams(uriPaths)
			if err != nil {
				fmt.Println(params)
			}
		}
	}
}


