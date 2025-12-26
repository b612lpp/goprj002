package meter

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/b612lpp/goprj002/application"
	"github.com/b612lpp/goprj002/domain"
	"github.com/b612lpp/goprj002/internal/delivery/http/exeptions"
	"github.com/b612lpp/goprj002/internal/middleware"
	"github.com/b612lpp/goprj002/repository"
)

type Meter struct {
	Uc *application.SubmitReading
}

func NewMeter(uc *application.SubmitReading) *Meter {
	return &Meter{Uc: uc}
}

func (m *Meter) GetValues(w http.ResponseWriter, r *http.Request) {
	//читаем роль из контекста. От него будет ветвится сценарий, а так же формироваться идентификатор для записи объекта показаний
	uid := r.Context().Value(middleware.OwnerId{}).(string)
	//создаём экземпляр объекта показаний, с предопределенным типом счетчика
	mr := domain.NewGasReading(uid)
	//Забираем инт из JSON и аппендим в пустой массив нового экземпляра показаний
	tmpD, err := parseIncJ(r)
	if err != nil {
		w.WriteHeader(400)
		slog.Info("функция обработки данных вернула ", "ошибка ", err)
		return
	}
	mr.SetValue(tmpD)
	slog.Info("получены данные", "тип счетчика газ. показания", tmpD)

	err = m.Uc.Execute(mr)
	switch {
	case err == nil:
		return
	case errors.Is(err, repository.ErrEmptyData):
		slog.Info("нет предыдущих значений")
		w.WriteHeader(204)
	case errors.Is(err, application.ErrValueValidation):
		slog.Info("значение меньше предыдущего")
		w.WriteHeader(400)
	default:
		w.WriteHeader(500)
		slog.Error("неизвестная ошибка")

	}

}

func parseIncJ(r *http.Request) ([]int, error) {
	type t struct {
		I []int `json:"value"`
	}
	q := t{}
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {

		return nil, exeptions.ErrParseData
	} else {

		return q.I, nil
	}

}
