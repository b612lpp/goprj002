package middleware

import (
	"context"
	"fmt"
	"net/http"
)

type UserId struct{}

func AuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello from private area")
		ActualCtx := r.Context()
		//Проверяем факт аутентификации. Сейчас в заголовке
		if h := r.Header["Auth"]; h[0] != "true" {
			fmt.Println("хедер Auth пришел в MW", r.Header["Auth"])
			w.WriteHeader(401)
			return
		} else {
			fmt.Println("хедер Auth пришел в MW", r.Header["Auth"])
			ctx := context.WithValue(ActualCtx, UserId{}, "me")

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
