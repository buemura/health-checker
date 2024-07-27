package entity

import "time"

type Endpoint struct {
	ID             string
	Name           string
	Url            string
	Status         string
	CheckFrequency int
	LastChecked    *time.Time
	NotifyTo       string
}
