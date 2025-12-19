package auth

import "net/http"

type Auth struct {
	Name string
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Registry(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (a *Auth) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
