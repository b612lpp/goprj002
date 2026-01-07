package meter

import (
	"errors"
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

// Структура для получения пользовательских данных
type enValues struct {
	Day   int `json:"day"`
	Night int `json:"night"`
}

func (me *EnMeterHandler) GetEnValues(w http.ResponseWriter, r *http.Request) {

	//Вычитываем пользователя из аутентификатора и создаем экземпляр доменного объекта
	uid := r.Context().Value(middleware.OwnerId{}).(string)
	emr := domain.NewEnReading(uid)

	//Инициализируем структуру для получения пользовательских данных и передаём её адрес в парсер
	t := enValues{}
	if err := parseIncJ(r, &t); err != nil {
		slog.Error("Ошибка обработки входящих данных")
		w.WriteHeader(400)
		return
	}

	//После успешного парсинга заполняем модель
	if err := emr.SetValue([]int{t.Day, t.Night}); err != nil {
		w.WriteHeader(500)
		slog.Info("ошибка процессинга данных", "подробности", err)
		return
	}

	//Передаем заполненный объект в юз кейс
	slog.Info("данные переданы на обработку", "скоуп значений", t)
	err := me.Uc.Execute(emr)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	if errors.Is(err, domain.ErrNewValueLessThanPrev) {
		slog.Info("значение меньше предыдущего")
		w.WriteHeader(400)
		return
	}
	if errors.Is(err, domain.ErrValueLessThanZero) {
		slog.Info("входящее значение меньше нуля")
		w.WriteHeader(400)
		return
	}
	if errors.Is(err, domain.ErrEmptyValues) {
		slog.Info("входящее значение не могут быть пустыми")
		w.WriteHeader(400)
		return
	}

	slog.Error("неизвестная ошибка", "err", err)
	w.WriteHeader(500)
}
