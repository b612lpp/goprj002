package meter

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/b612lpp/goprj002/application"
	"github.com/b612lpp/goprj002/domain"
	"github.com/b612lpp/goprj002/internal/middleware"
)

type GasMeterHandler struct {
	Uc application.SubmitReadingGas
}

func NewGasMeterHandler(uc application.SubmitReadingGas) *GasMeterHandler {
	return &GasMeterHandler{Uc: uc}
}

type gasValues struct {
	I int `json:"value"`
}

func (m *GasMeterHandler) GetGasValues(w http.ResponseWriter, r *http.Request) {

	//Вычитываем пользователя из аутентификатора и создаем экземпляр доменного объекта
	uid := r.Context().Value(middleware.OwnerId{}).(string)
	//Инициализируем структуру для получения пользовательских данных и передаём её адрес в парсер
	t := gasValues{}
	if err := parseIncJ(r, &t); err != nil {
		slog.Error("Ошибка обработки входящих данных")
		w.WriteHeader(400)
		return
	}

	//Передаем заполненный объект в юз кейс входящие значения и ИДпользователя
	slog.Info("данные переданы на обработку", "скоуп значений", t)
	err := m.Uc.Execute(uid, []int{t.I})
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
