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
func (mr *MeterReading) IsValid(p []int) error {

	if len(mr.Values) == 0 || mr.Values == nil {
		return ErrEmptyValues
	}

	if len(p) == 0 {
		for i := range mr.Values {
			if mr.Values[i] < 0 {
				return ErrValueLessThanZero
			}

		}
		return nil
	}

	if len(p) != len(mr.Values) && len(p) != 0 {
		return ErrValuesMismatch
	}

	for i := range p {
		if mr.Values[i] < 0 {
			return ErrValueLessThanZero
		}
		if p[i] > mr.Values[i] {
			return ErrNewValueLessThanPrev
		}

	}

	return nil
}

func NewGasReading(owner string) MeterReading {
	return MeterReading{ownerId: owner, meterType: "_Gas", CreatedAt: time.Now()}
}
func NewEnReading(owner string) MeterReading {

	return MeterReading{ownerId: owner, meterType: "_Electro", CreatedAt: time.Now()}
}
