package application

import (
	"time"

	"github.com/b612lpp/goprj002/domain"
)

type EventFabric struct {
}

type EventFormer interface {
	MakeEvent(mr domain.MeterReading) domain.EventuallyAppliedData
}

func (ef *EventFabric) MakeEvent(mr domain.MeterReading) domain.EventuallyAppliedData {
	return domain.EventuallyAppliedData{Owner: mr.GetOwnerID(), MeterType: mr.GetMEterType(), Values: mr.GetValues(), CreatedAt: time.Now()}
}

func NewEventFabric() EventFormer {
	return &EventFabric{}
}
