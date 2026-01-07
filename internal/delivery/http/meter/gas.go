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
	I []int `json:"value"`
}

func (m *GasMeterHandler) GetGasValues(w http.ResponseWriter, r *http.Request) {
	//Инициализируем пустую структуру для пользовательских данных
	t := gasValues{}

	//читаем роль из контекста. От него будет ветвится сценарий, а так же формироваться идентификатор для записи объекта показаний
	uid := r.Context().Value(middleware.OwnerId{}).(string)
	//создаём экземпляр объекта показаний, с предопределенным типом счетчика
	mr := domain.NewGasReading(uid)
	//Забираем инт из JSON и аппендим в пустой массив нового экземпляра показаний
	err := parseIncJ(r, &t)
	if err != nil {
		w.WriteHeader(400)
		slog.Info("функция обработки данных вернула ", "ошибка ", err)
		return
	}
	mr.SetValue(t.I)
	slog.Info("получены данные", "тип счетчика газ. показания", t.I)

	err = m.Uc.Execute(mr)
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
