package api

type Category struct {
	Id          int           `json:"id,omitempty"`
	Subject     int           `json:"subject,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
}
