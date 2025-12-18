package health

import (
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (hh *HealthHandler) ResponsOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

}
