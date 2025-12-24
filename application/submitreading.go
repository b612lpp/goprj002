package application

import (
	"fmt"

	"github.com/b612lpp/goprj002/domain"
	"github.com/b612lpp/goprj002/repository"
)

type AllUseCases struct {
}

type SubmitReading struct {
	R repository.Repo
}

func NewSubmitReading(r repository.Repo) *SubmitReading {
	return &SubmitReading{R: r}
}

func (s *SubmitReading) Execute(mr domain.MeterReading) error {
	if z, err := s.R.GetLast("user", "gas"); err != nil {

		return err
	} else {
		fmt.Println(z)
		return nil
	}

	//q:=s.R.GetLast(u)

}
