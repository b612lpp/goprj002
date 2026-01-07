package application

import (
	"log/slog"

	"github.com/b612lpp/goprj002/domain"
	"github.com/b612lpp/goprj002/repository"
)

type AllUseCases struct {
}

type SubmitReadingGas struct {
	R repository.Repo
}

func NewSubmitReadingGas(r repository.Repo) *SubmitReadingGas {
	return &SubmitReadingGas{R: r}
}

func (s *SubmitReadingGas) Execute(u string, v []int) error {
	gmr := domain.NewGasReading(u)
	gl, err := s.R.GetLast(u, gmr.GetMEterType())
	if err != nil && err != repository.ErrEmptyData {
		slog.Error("ошибка получения предыдущих показаний", "owner", gmr.GetOwnerID(), "err", err)
		return err
	}

	if err = gmr.Apply(gl.GetValues(), v); err != nil {
		return err
	}

	s.R.Save(gmr)
	slog.Info("данные добавлены в бд", "owner", gmr.GetOwnerID(), "new_values", gmr.GetValues(), "previous", gl.GetValues())
	return nil
}
