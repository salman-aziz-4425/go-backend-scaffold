package models

type Video struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorId int    `json:"authorId"`
	Author   User   `gorm:"foreignKey:AuthorId"`
}
