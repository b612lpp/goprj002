package domain

import "time"

type EventuallyAppliedData struct {
	owner     string
	meterType string
	values    []int
	createdAt time.Time
}
