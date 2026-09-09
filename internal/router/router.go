package router

import (
	"net/http"
	"strings"
)

type MiddlewareHandler func(http.Handler) http.Handler

type Router struct {
	*http.ServeMux
	middlewares []MiddlewareHandler
}

func NewRouter() *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var h http.Handler = r.ServeMux
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		h = r.middlewares[i](h)
	}
	h.ServeHTTP(w, req)
}

func (r *Router) Use(handler MiddlewareHandler) {
	r.middlewares = append(r.middlewares, handler)
}

func (r *Router) HandleGroup(prefix string, group http.Handler) {
	prefix = strings.TrimSuffix(prefix, "/")
	r.Handle(prefix+"/", http.StripPrefix(prefix, group))
}
