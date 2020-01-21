package api

type Subject struct {
	Id         int    `json:"id,omitempty"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"name,omitempty"`
}
