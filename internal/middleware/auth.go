package middleware

import (
	"context"
	"log/slog"
	"net/http"
)

type UserInfo struct {
	Id, Role string
}

func AuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if userData, err := CheckHeaders(r.Header); err != nil {
			w.WriteHeader(401)
			slog.Info("Ошибка аутентификации")
			return
		} else {
			ActualCtx := r.Context()
			Ctx := context.WithValue(ActualCtx, UserInfo{}, userData)
			next.ServeHTTP(w, r.WithContext(Ctx))
		}

	})
}

func CheckHeaders(h http.Header) (UserInfo, error) {
	id, i := h["Auth"]
	role, r := h["Role"]
	if i == true && r == true {
		return UserInfo{Id: id[0], Role: role[0]}, nil
	}
	return UserInfo{}, ErrBadCreds
}
