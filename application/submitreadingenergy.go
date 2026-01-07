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

func (s *SubmitReadingEn) Execute(mr domain.MeterReading) error {
	//Получаем предыдущее значение и ошибку. От этого действуем
	gl, err := s.R.GetLast(mr.GetOwnerID(), mr.GetMEterType())

	if err != nil && err != repository.ErrEmptyData {
		slog.Error("ошибка получения предыдущих показаний", "owner", mr.GetOwnerID(), "err", err)
		return err
	}

	if err := mr.IsValid(gl.Values); err != nil {
		return err
	}
	s.R.Save(mr)
	slog.Info("данные добавлены в бд", "owner", mr.GetOwnerID(), "new_values", mr.Values, "previous", gl.Values)
	return nil
}
