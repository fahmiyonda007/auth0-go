// controllers/books.go

package controllers

import (
	"net/http"
	"strconv"

	auth "github.com/fahmiyonda007/go-gin-gorm/controllers/auth"
	metadata "github.com/fahmiyonda007/go-gin-gorm/controllers/pagination"
	"github.com/fahmiyonda007/go-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

var readPermission = "read:book"
var createPermission = "create:book"
var updatePermission = "update:book"
var deletePermission = "delete:book"

// GET /books
// Get all books
func Books(c *gin.Context) {
	token, err := auth.Authenticate(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if !auth.HasPermission(token, readPermission) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		return
	}

	//url.domain?page=1&length=10
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	length, err := strconv.ParseInt(c.DefaultQuery("length", "10"), 10, 64)

	var books []models.Book
	models.DB.Find(&books).Limit(metadata.Limit(int(length))).Offset(metadata.Offset(int(page), int(length)))

	c.JSON(http.StatusOK, gin.H{
		"metadata": metadata.CalculateMetadata(len(books), int(page), int(length)),
		"data":     books,
	})
}

// GET /books/:id
// Find a book
func Book(c *gin.Context) { // Get model if exist
	token, err := auth.Authenticate(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if !auth.HasPermission(token, readPermission) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		return
	}

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	token, err := auth.Authenticate(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if !auth.HasPermission(token, createPermission) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		return
	}

	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})

}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	token, err := auth.Authenticate(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if !auth.HasPermission(token, updatePermission) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		return
	}

	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Record not found!",
		})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	token, err := auth.Authenticate(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if !auth.HasPermission(token, deletePermission) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		return
	}

	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Record not found!",
		})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}