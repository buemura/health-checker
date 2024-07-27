package dto

type CreateEndpointIn struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	CheckFrequency int    `json:"check_frequency"`
	NotifyTo       string `json:"notify_to"`
}
