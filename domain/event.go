package domain

import "time"

type EventuallyAppliedData struct {
	Owner     string
	MeterType string
	Values    []int
	CreatedAt time.Time
}
