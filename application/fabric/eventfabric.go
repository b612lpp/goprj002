package fabric

import (
	"github.com/b612lpp/goprj002/domain"
)

type EventFabric struct {
}

type EventFormer interface {
	MakeEvent(mr domain.MeterReading) domain.EventuallyAppliedData
}

func (ef *EventFabric) MakeEvent(mr domain.MeterReading) domain.EventuallyAppliedData {
	return domain.EventuallyAppliedData{Owner: mr.GetOwnerID(), MeterType: mr.GetMEterType(), Values: mr.GetValues(), CreatedAt: mr.GetCreatedAt()}
}

func NewEventFabric() EventFormer {
	return &EventFabric{}
}
