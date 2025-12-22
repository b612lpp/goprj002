package meter

import (
	"fmt"
	"net/http"

	"github.com/b612lpp/goprj002/internal/middleware"
)

type Meter struct {
	Tmp string
}

func NewMeter() *Meter {
	return &Meter{Tmp: "Временная запись"}
}

func (m *Meter) TmpMeter(w http.ResponseWriter, r *http.Request) {
	ActualCtx := r.Context()
	fmt.Println("Запрос от пользователя ", ActualCtx.Value(middleware.UserInfo{}))
}
