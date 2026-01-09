package application

import (
	"log/slog"

	"github.com/b612lpp/goprj002/application/fabric"
	"github.com/b612lpp/goprj002/domain"
	"github.com/b612lpp/goprj002/repository"
)

type SubmitReadingEn struct {
	R repository.ReadingStorage
	F fabric.EventFormer
}

func NewSubmitReadingEn(r repository.ReadingStorage, f fabric.EventFormer) *SubmitReadingEn {
	return &SubmitReadingEn{R: r, F: f}
}

func (s *SubmitReadingEn) Execute(u string, v []int) error {

	emr := domain.NewEnReading(u)
	gl, err := s.R.GetLast(u, emr.GetMEterType())
	if err != nil && err != repository.ErrEmptyData {
		slog.Error("ошибка получения предыдущих показаний", "owner", emr.GetOwnerID(), "err", err)
		return err
	}

	err = emr.Apply(gl.GetValues(), v)
	if err != nil {
		return err
	}

	err = s.R.Save(emr)
	if err != nil {
		slog.Info("ошибка записи в БД")
		return err
	}
	slog.Info("данные добавлены в бд", "owner", emr.GetOwnerID(), "new_values", emr.GetValues(), "previous", gl.GetValues())

	if err = s.R.AddEvent(s.F.MakeEvent(emr)); err != nil {
		slog.Info("ошибка сохранения события")
		return err
	}

	return nil
}
