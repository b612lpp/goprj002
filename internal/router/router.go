package router

import (
	"net/http"
)

type Router interface {
	Handler() http.Handler
}

type MyRouter struct {
	Mux *http.ServeMux
}

func NewMyRouter() *MyRouter {
	q := http.NewServeMux()
	return &MyRouter{Mux: q}
}

func (mr *MyRouter) AddRout(p string, h http.HandlerFunc) {
	mr.Mux.Handle(p, h)
}

func (mr *MyRouter) Handler() http.Handler {
	return mr.Mux
}
