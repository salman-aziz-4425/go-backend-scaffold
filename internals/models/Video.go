package models

type Video struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorId string `json:"authorId"`
}
