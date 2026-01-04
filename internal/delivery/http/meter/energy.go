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
	emr.SetValue([]int{t.Day, t.Night})

	//Передаем заполненный объект в юз кейс
	slog.Info("данные переданы на обработку", "скоуп значений", t)
	err := me.Uc.Execute(emr)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	if errors.Is(err, application.ErrValueValidation) {
		slog.Info("значение меньше предыдущего")
		w.WriteHeader(400)
		return
	}

	slog.Error("неизвестная ошибка", "err", err)
	w.WriteHeader(500)
}
