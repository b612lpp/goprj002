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
	MainMux, PrivateMux, PublicMux *http.ServeMux
}

func NewMyRouter() *MyRouter {
	return &MyRouter{
		MainMux:    http.NewServeMux(),
		PrivateMux: http.NewServeMux(),
		PublicMux:  http.NewServeMux(),
	}
}

// Добавить публичный маршрут в mux
func (mr *MyRouter) AddPublicRout(p string, h http.HandlerFunc) {
	mr.PublicMux.HandleFunc(p, h)
}

// Добавить приватный маршрут в mux
func (mr *MyRouter) AddPrivateRout(p string, h http.HandlerFunc) {
	mr.PrivateMux.HandleFunc(p, h)
}

//Собираем MainMux

func (mr *MyRouter) CompilemmMux() {
	mr.MainMux.Handle("/private/", middleware.AuthMW(mr.PrivateMux))
	mr.MainMux.Handle("/public/", mr.PublicMux)
}

// Самое вкусненькое. Возвращаем обернутый mux
func (mr *MyRouter) Handler() http.Handler {
	return middleware.Logging(mr.MainMux)
}
