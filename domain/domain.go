package domain

import "time"

type AllMeters interface {
	GetName() string
	New() []int
}

type GasMeter struct {
	Title string
	Value int
}

type ElMeter struct {
	Title            string
	Day, Night, Summ int
}

type WaterMeter struct {
	Title     string
	cool, hot int
}

type MeterReading struct {
	ownerId   string
	meterType string
	CreatedAt time.Time
	Values    []int `json:"value"`
}

func NewMeterReading(o, m string) MeterReading {
	return MeterReading{ownerId: o, meterType: m}
}

func (mr *MeterReading) GetOwnerID() string {
	return mr.ownerId
}

func (mr *MeterReading) GetMEterType() string {
	return mr.meterType
}

func NewGasReading(owner string) MeterReading {
	return MeterReading{ownerId: owner, meterType: "_Gas"}
}
func NewElectricityReading(owner string, day, night int) MeterReading {

	return MeterReading{ownerId: owner, meterType: "_Electro", Values: []int{day, night}}
}
