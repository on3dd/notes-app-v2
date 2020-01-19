package api

type Teacher struct {
	Id        int    `json:"id,omitempty"`
	SubjectId int    `json:"subject_id,omitempty"`
	Name      string `json:"name,omitempty"`
}
