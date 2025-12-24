package middleware

import (
	"context"
	"log/slog"
	"net/http"
)

type OwnerId struct {
}

type OwnerRole struct {
}

type UserInfo struct {
	Id, Role, Status string
}

func AuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if userData, err := CheckHeaders(r.Header); err != nil {
			w.WriteHeader(401)
			slog.Info("Ошибка аутентификации")
			return
		} else {
			ctx := r.Context()
			ctx = context.WithValue(ctx, OwnerId{}, userData.Id)
			ctx = context.WithValue(ctx, OwnerRole{}, userData.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

	})
}

func CheckHeaders(h http.Header) (UserInfo, error) {
	authStatus, i := h["Auth"]
	id, _ := h["Login"]
	role, r := h["Role"]
	if i == true && r == true {

		return UserInfo{Status: authStatus[0], Role: role[0], Id: id[0]}, nil
	}
	return UserInfo{}, ErrBadCreds
}
