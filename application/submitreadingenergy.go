package application

import (
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

func (s *SubmitReadingEn) Execute(u string, v []int) error {

	emr := domain.NewEnReading(u)
	gl, err := s.R.GetLast(u, emr.GetMEterType())
	if err != nil && err != repository.ErrEmptyData {
		slog.Error("ошибка получения предыдущих показаний", "owner", emr.GetOwnerID(), "err", err)
		return err
	}

	if err = emr.Apply(gl.GetValues(), v); err != nil {
		return err
	}

	s.R.Save(emr)
	slog.Info("данные добавлены в бд", "owner", emr.GetOwnerID(), "new_values", emr.GetValues(), "previous", gl.GetValues())
	return nil
}
