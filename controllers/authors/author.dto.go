package controllers

import (
	"github.com/fahmiyonda007/go-gin-gorm/models"
)

type CreateAuthorInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAuthorInput struct {
	Name string `json:"name"`
}

type AuthorOutput struct {
	ID   uint
	Name string
	Book []models.Book
}
