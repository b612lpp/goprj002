package meter

import (
	"encoding/json"
	"fmt"
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

	uid := r.Context().Value(middleware.OwnerId{}).(string)

	//ur := ActualCtx.Value(middleware.OwnerRole{})
	mr := domain.NewGasReading(uid)
	//Забираем инт из JSON и аппендим в пустой массив нового экземпляра показаний
	mr.Values = append(mr.Values, parseIncJ(r))

	fmt.Println(mr)

	if err := m.Uc.Execute(mr); err != nil {
		w.WriteHeader(400)
		return
	}
}

func parseIncJ(r *http.Request) int {
	type t struct {
		I int `json:"value"`
	}
	q := t{}
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		fmt.Println("не успех")
		return -1
	} else {
		fmt.Println("успех парсинга")
		fmt.Println(q.I)
		return q.I
	}

}
