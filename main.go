package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type userid struct{}
type debug struct{}

func main() {

	h := http.NewServeMux()
	h.HandleFunc("/", homePage)
	wh := xDeburRead(mwCont(h))
	log.Fatal(http.ListenAndServe(":8081", wh))
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Context().Value(userid{}))

	if r.Context().Value(debug{}) == true {
		fmt.Println("Есть дебаг")
		return
	}
	fmt.Println("нет дебага")

}

func mwCont(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctxOgig := r.Context() //читаем текущий контекст

		ctx := context.WithValue(ctxOgig, userid{}, 123)
		newR := r.WithContext(ctx)
		next.ServeHTTP(w, newR)

	})
}

func xDeburRead(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		q := r.Header
		uniqCtx := r.Context()
		if v := q["X-Debug"]; v[0] != "" {
			ctx := context.WithValue(uniqCtx, debug{}, true)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		ctx := context.WithValue(uniqCtx, debug{}, false)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
