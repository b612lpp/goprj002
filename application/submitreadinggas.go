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
		slog.Error("ошибка получения предыдущих показаний", "owner", mr.GetOwnerID(), "err", err)
		return err
	}

	if mr.Validate() != true {
		slog.Info("Полученные данные меньше 0")
		return ErrLowerZero
	}

	if len(gl.Values) == 0 && err == repository.ErrEmptyData {
		if err := s.R.Save(mr); err != nil {
			return err
		}
		slog.Info("новые данные записаны в бд", "owner", mr.GetOwnerID(), "value", mr.Values)
		return nil
	}

	if res := mr.IsValidComparedTo(gl.Values); res != true {
		return ErrValueValidation
	}
	s.R.Save(mr)
	slog.Info("данные добавлены в бд", "owner", mr.GetOwnerID(), "new_values", mr.Values, "previous", gl.Values)
	return nil
}
