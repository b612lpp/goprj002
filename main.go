package main

import (
	"fmt"
	"net/http"
)

type Router interface {
	Handle() http.Handler
}

type MyRouter struct {
	MainMux, PrivateMux, PublicMux *http.ServeMux
}

func NewMyrouter() *MyRouter {
	return &MyRouter{MainMux: http.NewServeMux(), PrivateMux: http.NewServeMux(), PublicMux: http.NewServeMux()}

}

func (mr *MyRouter) AddPrivateRoute(p string, h http.HandlerFunc) {
	mr.PrivateMux.Handle(p, h)
}

func (mr *MyRouter) AddPublicRoute(p string, h http.HandlerFunc) {
	mr.PublicMux.Handle(p, h)
}

func PrivateF(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from Privat area")
}
func PublicF(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from Public area")
}

func (mr *MyRouter) Handle() http.Handler {
	return mr.MainMux
}

func GeneralMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		fmt.Println("hello from general MW")
		next.ServeHTTP(w, r)
	})
}

func PrivatMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello from privat MW")
		next.ServeHTTP(w, r)
	})
}
func PublicMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello from public MW")

		next.ServeHTTP(w, r)
	})
}

func main() {
	h := NewMyrouter()
	h.AddPrivateRoute("/private/p", PrivateF)
	h.AddPublicRoute("/public/p", PublicF)

	h.MainMux.Handle("/private/", PrivatMW(h.PrivateMux))
	h.MainMux.Handle("/public/", PublicMW(h.PublicMux))
	http.ListenAndServe(":8081", GeneralMW(h.MainMux))
}
