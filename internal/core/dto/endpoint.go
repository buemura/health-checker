package dto

import "time"

type CreateEndpointIn struct {
	Name           string
	Url            string
	Status         string
	CheckFrequency int
	LastChecked    time.Time
	NotifyTo       string
}
