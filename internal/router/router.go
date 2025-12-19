package router

import (
	"net/http"

	"github.com/b612lpp/goprj002/internal/middleware"
)

// наш кастомный роутер должен соответствовать
type Router interface {
	Handler() http.Handler
}

// Роутер. По сути обертка над mux. Позволит масштабироваться, избежать прямых записей в Mux без участия роутера
type MyRouter struct {
	Mux http.ServeMux
}

func NewMyRouter() *MyRouter {
	return &MyRouter{Mux: *http.NewServeMux()}
}

// Добавить маршрут в mux
func (mr *MyRouter) AddRout(p string, h http.HandlerFunc) {
	mr.Mux.HandleFunc(p, h)
}

// Самое вкусненькое. Возвращаем обернутый mux
func (mr *MyRouter) Handler() http.Handler {
	return middleware.Logging(&mr.Mux)
}
