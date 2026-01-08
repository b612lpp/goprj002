package repository

import (
	"time"

	"github.com/b612lpp/goprj002/domain"
)

type EventDb struct {
	title string
	event []domain.EventuallyAppliedData
}

func NewEventDb() EventStore {

	return &EventDb{title: "InMemory Evennt DB"}
}

func (ed *EventDb) AddEvent(v domain.EventuallyAppliedData) error {
	ed.event = append(ed.event, v)
	return nil
}

func (ed *EventDb) ReadEvens(owner string, mtype string, from, to time.Time) []domain.EventuallyAppliedData {
	return nil
}

func (edb *EventDb) GetTitle() string {
	return edb.title
}
