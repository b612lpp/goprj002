package application

import (
	"fmt"
	"log/slog"

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

	gl, err := s.R.GetLast(mr.GetOwnerID(), mr.GetMEterType())
	if err != nil && err != repository.ErrEmptyData {
		return err
	}
	if len(gl.Values) > 0 && mr.Values[0] >= gl.Values[0] {
		s.R.Save(mr)
		slog.Info("данные добавлены в бд", "для идентификатора ", mr.GetOwnerID(), "новые показания ", mr.Values[0])
		fmt.Println(s.R.SelectAll())
		return nil

	}

	if len(gl.Values) == 0 {
		s.R.Save(mr)
		slog.Info("новые данные записаны в бд")

		return nil
	}
	slog.Info("неизвестная ошибка", "сервер вернул", ErrUnknown)
	return ErrUnknown

}
