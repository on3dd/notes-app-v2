package api

type SubjectCard struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	PostsCount   int    `json:"posts_count"`
}
