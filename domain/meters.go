package domain

type GasMeter struct {
	Ttitle    string
	ValueType string
	Values    [1]int
}

type EnergyMeter struct {
	Ttitle    string
	ValueType string
	Values    [2]int
}

type WaterMeter struct {
	Ttitle    string
	ValueType string
	Values    [2]int
}

type Meters interface {
	GasMeter | EnergyMeter | WaterMeter
}
