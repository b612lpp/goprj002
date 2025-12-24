package repository

import (
	"github.com/b612lpp/goprj002/domain"
)

type IMDB struct {
	Data map[string][]domain.MeterReading
}

func NemIMDB() Repo {
	return &IMDB{Data: make(map[string][]domain.MeterReading)}
}

func (db *IMDB) Save(mr domain.MeterReading) error {
	key := mr.GetOwnerID() + mr.GetMEterType()
	db.Data[key] = append(db.Data[key], mr)
	return nil
}

func (db *IMDB) GetLast(u, t string) (domain.MeterReading, error) {
	q := len(db.Data[u+t]) - 1
	if qq, ok := db.Data[u+t]; ok != false {
		return qq[q], nil
	}
	return domain.MeterReading{}, ErrEmptyData
}
