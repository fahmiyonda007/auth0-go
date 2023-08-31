package controllers

import (
	"github.com/fahmiyonda007/go-gin-gorm/models"
)

type CreateBookInput struct {
	Title    string `json:"title" binding:"required"`
	AuthorId uint   `json:"authorId" binding:"required"`
}

type UpdateBookInput struct {
	Title    string `json:"title"`
	AuthorId uint   `json:"authorId"`
}

type BookOutput struct {
	ID     uint
	Title  string
	Author models.Author
}
