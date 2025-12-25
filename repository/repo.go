package repository

import (
	"github.com/b612lpp/goprj002/domain"
)

type Repo interface {
	GetLast(string, string) (domain.MeterReading, error)

	Save(reading domain.MeterReading) error
	SelectAll() *IMDB
}
