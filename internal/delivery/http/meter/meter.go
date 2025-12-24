package meter

import (
	"net/http"

	"github.com/b612lpp/goprj002/application"
	"github.com/b612lpp/goprj002/domain"
	"github.com/b612lpp/goprj002/internal/middleware"
)

type Meter struct {
	Uc *application.SubmitReading
}

func NewMeter(uc *application.SubmitReading) *Meter {
	return &Meter{Uc: uc}
}

func (m *Meter) GetValues(w http.ResponseWriter, r *http.Request) {

	ActualCtx := r.Context()

	uid := ActualCtx.Value(middleware.OwnerId{})
	ur := ActualCtx.Value(middleware.OwnerRole{})
	q := domain.NewMeterReading(uid.(string), ur.(string))
	if err := m.Uc.Execute(q); err != nil {
		w.WriteHeader(400)
		return
	}
}
