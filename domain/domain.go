package domain

import "time"

//Описание модели для работы с входящими данными показаний счетчиков клиента
type MeterReading struct {
	ownerId   string
	meterType string
	counts    int //количество показателей в счетчике
	values    []int
	createdAt time.Time
}

func (mr *MeterReading) GetOwnerID() string {
	return mr.ownerId
}

func (mr *MeterReading) GetMEterType() string {
	return mr.meterType
}
func (mr *MeterReading) GetValues() []int {
	cp := make([]int, len(mr.values))
	copy(cp, mr.values)
	return cp
}

//Сравниваем полученные значения с предыдущими, если ОК то заполняем агрегат
func (mr *MeterReading) Apply(p, v []int) (EventuallyAppliedData, error) {

	if len(v) != mr.counts {
		return EventuallyAppliedData{}, ErrValuesTypeMismatch
	}

	for i := range v {
		if v[i] < 0 {
			return EventuallyAppliedData{}, ErrValueLessThanZero
		}

	}
	if len(v) == len(p) {
		for i := range v {
			if v[i] < p[i] {
				return EventuallyAppliedData{}, ErrNewValueLessThanPrev
			}
		}
	}

	mr.values = v

	return EventuallyAppliedData{owner: mr.ownerId, meterType: mr.meterType, values: mr.values, createdAt: time.Now()}, nil
}

func NewGasReading(owner string) MeterReading {
	return MeterReading{ownerId: owner, meterType: "_Gas", createdAt: time.Now(), counts: 1}
}
func NewEnReading(owner string) MeterReading {

	return MeterReading{ownerId: owner, meterType: "_Electro", createdAt: time.Now(), counts: 2}
}
