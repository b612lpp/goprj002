package repository

import (
	"github.com/b612lpp/goprj002/domain"
)

type IMDB struct {
	Title string
	Data  map[string][]domain.MeterReading
}

func NemIMDB() Repo {
	return &IMDB{Title: "IMDB", Data: make(map[string][]domain.MeterReading)}
}
func (db *IMDB) GetTitle() string {
	return db.Title
}
func (db *IMDB) Save(mr domain.MeterReading) error {
	key := mr.GetOwnerID() + mr.GetMEterType()
	db.Data[key] = append(db.Data[key], mr)
	return nil
}

// Возвращает последнюю запись по ключу юзер(u)+тип счетчика(t)
func (db *IMDB) GetLast(u, t string) (domain.MeterReading, error) {
	tmpLast := len(db.Data[u+t]) - 1
	if tmpLast >= 0 {
		return db.Data[u+t][tmpLast], nil
	}

	return domain.MeterReading{}, ErrEmptyData
}

func (db *IMDB) SelectAll() *IMDB {
	return db
}
