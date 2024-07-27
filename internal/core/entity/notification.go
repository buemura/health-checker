package entity

import "time"

type Notification struct {
	ID          string
	EndpointID  string
	Destination string
	CreatedAt   time.Time
}
