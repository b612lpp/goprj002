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
	if len(v) < 1 {
		return ErrValueToAdd
	}
	mr.Values = v
	return nil
}

//Сравниваем значение текущего значения с предыдущего. р - заданный массив дляпроверки
func (mr *MeterReading) IsValidComparedTo(p []int) bool {
	state := false
	if len(p) == len(mr.Values) {
		for i := range p {
			if p[i] <= mr.Values[i] {
				state = true
			} else {
				return false
			}

		}

	}
	return state
}

func (mr *MeterReading) Validate() bool {
	state := false
	for i := range mr.Values {
		if mr.Values[i] < 0 {
			return false
		}
		state = true
	}
	return state
}

func NewGasReading(owner string) MeterReading {
	return MeterReading{ownerId: owner, meterType: "_Gas", CreatedAt: time.Now()}
}
func NewEnReading(owner string) MeterReading {

	return MeterReading{ownerId: owner, meterType: "_Electro", CreatedAt: time.Now()}
}
