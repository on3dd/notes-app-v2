package api

type Teacher struct {
	Id         int    `json:"id,omitempty"`
	CategoryId int    `json:"category_id,omitempty"`
	Name       string `json:"name,omitempty"`
}
