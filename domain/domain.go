package domain

import "time"

type MeterReading struct {
	ownerId   string
	meterType string
	CreatedAt time.Time
	Values    []int `json:"value"`
}

func (mr *MeterReading) GetOwnerID() string {
	return mr.ownerId
}

func (mr *MeterReading) GetMEterType() string {
	return mr.meterType
}

func (mr *MeterReading) SetValue(v []int) error {
	if len(v) != 1 {
		return ErrValueToAdd
	}
	mr.Values = v
	return nil
}

func NewGasReading(owner string) MeterReading {
	return MeterReading{ownerId: owner, meterType: "_Gas", CreatedAt: time.Now()}
}
func NewEnReading(owner string) MeterReading {

	return MeterReading{ownerId: owner, meterType: "_Electro", CreatedAt: time.Now()}
}
