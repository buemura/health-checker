package dto

type CreateNotificationIn struct {
	EndpointID  string `json:"endpoint_id"`
	Destination string `json:"destination"`
}
