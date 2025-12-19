package router

import (
	"net/http"

	"github.com/b612lpp/goprj002/internal/middleware"
)

type Router interface {
	Handler() http.Handler
}

type MyRouter struct {
	Mux http.ServeMux
}

func NewMyRouter() *MyRouter {
	return &MyRouter{Mux: *http.NewServeMux()}
}

func (mr *MyRouter) AddRout(p string, h http.HandlerFunc) {
	mr.Mux.HandleFunc(p, h)
}

func (mr *MyRouter) Handler() http.Handler {
	return middleware.Logging(&mr.Mux)
}
