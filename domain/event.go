package domain

import "time"

type EventuallyAppliedData struct {
	owner     OwnerID
	meterType string
	values    []int
	createdAt time.Time
}
