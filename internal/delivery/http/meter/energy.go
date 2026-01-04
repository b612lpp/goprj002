package meter

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/b612lpp/goprj002/application"
	"github.com/b612lpp/goprj002/domain"
	"github.com/b612lpp/goprj002/internal/middleware"
)

type EnMeterHandler struct {
	Uc application.SubmitReadingEn
}

func NewEnMeterHandler(uc application.SubmitReadingEn) *EnMeterHandler {
	return &EnMeterHandler{Uc: uc}
}

type enValues struct {
	Day   int `json:"day"`
	Night int `json:"night"`
}

func (me *EnMeterHandler) GetEnValues(w http.ResponseWriter, r *http.Request) {
	v := enValues{}
	fmt.Println("дошли до хэндлера")
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		slog.Error("ошибка обработки данных")
	}
	uid := r.Context().Value(middleware.OwnerId{}).(string)
	emr := domain.NewEnReading(uid, v.Day, v.Night)
	me.Uc.Execute(emr)
}
