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

func (s *SubmitReadingGas) Execute(mr domain.MeterReading) error {
	gl, err := s.R.GetLast(mr.GetOwnerID(), mr.GetMEterType())
	if err != nil && err != repository.ErrEmptyData {
		return err
	}

	if len(mr.Values) == 0 {
		return ErrValueValidation
	}

	v := mr.Values[0]
	if v <= 0 {
		return ErrValueValidation
	}

	if len(gl.Values) == 0 {
		if err := s.R.Save(mr); err != nil {
			return err
		}
		slog.Info("новые данные записаны в бд", "owner", mr.GetOwnerID(), "value", v)
		return nil
	}

	last := gl.Values[0]
	if v < last {
		return ErrValueValidation
	}

	if err := s.R.Save(mr); err != nil {
		return err
	}
	slog.Info("данные добавлены в бд", "owner", mr.GetOwnerID(), "new_value", v, "previous", last)
	return nil
}
