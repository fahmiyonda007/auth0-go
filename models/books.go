package models

type Book struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	AuthorId uint   `json:"-" gorm:"foreignKey:AuthorId;references:ID"`
	Author   Author `json:"author"`
}
