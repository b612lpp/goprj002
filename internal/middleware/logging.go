package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// Обертка логирования для роутера. Версия до цепочки
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		slog.Info("Запрос принят", "Вызов метода", r.Method,
			"Вызов УРЛ", r.RequestURI,
			"вызов от ", r.RemoteAddr,
		)
		next.ServeHTTP(w, r)
		t := time.Since(start)
		slog.Info("Ответ отправлен",

			"время обработки ответ", t,
		)
	})

}
