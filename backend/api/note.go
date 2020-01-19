package api

type Note struct {
	Id          int    `json:"id,omitempty"`
	Author      int    `json:"author_id,omitempty"`
	CategoryId  int    `json:"category_id,omitempty"`
	TeacherId   int    `json:"teacher_id,omitempty"`
	PostedAt    string `json:"posted_at,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
}