package application

import (
	"fmt"
	"log/slog"

	"github.com/b612lpp/goprj002/domain"
	"github.com/b612lpp/goprj002/repository"
)

type SubmitReadingEn struct {
	R repository.Repo
}

func NewSubmitReadingEn(r repository.Repo) *SubmitReadingEn {
	return &SubmitReadingEn{R: r}
}

func (s *SubmitReadingEn) Execute(mr domain.MeterReading) error {
	if err := s.R.Save(mr); err != nil {
		fmt.Println("Ошибка")
		return err
	} else {

		slog.Info("данные приняты", "показания ", mr.Values)
		fmt.Println(s.R.SelectAll())
		return nil
	}

}
