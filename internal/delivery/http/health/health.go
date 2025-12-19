// Модуль-хендлер. Дает простые ответы о состоянии сервера
package health

import (
	"encoding/json"
	"net/http"
)

type Respons struct {
	Status string `json:"status"`
}

type HealthHandler struct {
}

// создаем пустой экземпляр
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// отвечаем ок на запрос
func (hh *HealthHandler) ResponsOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Respons{Status: "OK"})

}
