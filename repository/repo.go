package repository

import (
	"github.com/b612lpp/goprj002/domain"
)

type Repo interface {
	// Возвращает последнюю запись по ключу юзер(u)+тип счетчика(t)
	GetLast(string, string) (domain.MeterReading, error)

	Save(reading domain.MeterReading) error
	SelectAll() *IMDB
	GetTitle() string
}
