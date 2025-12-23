package application

import (
	"context"
	"fmt"

	"github.com/b612lpp/goprj002/domain"
	"github.com/b612lpp/goprj002/internal/middleware"
	"github.com/b612lpp/goprj002/repository"
)

type SubmitReading struct {
	R repository.Repo
}

func NewSubmitReading(r repository.Repo) *SubmitReading {
	return &SubmitReading{R: r}
}

func (s *SubmitReading) Execute(mr domain.MeterReading) error {
	u := context.Background().Value(middleware.UserInfo{})
	fmt.Println(u)
	//	q:=s.R.GetLast(u)
	return nil

}
