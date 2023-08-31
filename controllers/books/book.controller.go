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

//	@BasePath	/api/v1
//
// Books godoc
//
//	@Summary	get all books
//	@Schemes
//	@Description	get all books
//	@Tags			Books
//	@Param			page	query	int	false	"page"
//	@Param			length	query	int	false	"length"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		BookOutput
//	@Failure		401	{object}	handler.JSONResult
//	@Router			/books [get]
//	@Security		BearerAuth
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

	var count int64
	var books []models.Book
	models.DB.Find(&books).Count(&count).Limit(metadata.Limit(int(length))).Offset(metadata.Offset(int(page), int(length))).Find(&books)

	meta := metadata.CalculateMetadata(int(count), int(page), int(length))
	validate := metadata.ValidateFilter(meta, int(page), int(length))

	if validate != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": validate})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"metadata": meta,
		"data":     books,
	})
}

//	@BasePath	/api/v1
//
// Books/:id godoc
//
//	@Summary	find a book by id
//	@Schemes
//	@Description	find a book by id
//	@Tags			Books
//	@Param			id	path	int	false	"id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	BookOutput
//	@Router			/books/{id} [get]
//	@Security		BearerAuth
func Book(c *gin.Context) {
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

//	@BasePath	/api/v1
//
// Books godoc
//
//	@Summary	Create new book
//	@Schemes
//	@Description	Create new book
//	@Tags			Books
//	@Param			input	body	CreateBookInput	true	"Input"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	BookOutput
//	@Router			/books [post]
//	@Security		BearerAuth
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

//	@BasePath	/api/v1
//
// Books godoc
//
//	@Summary	Update a book
//	@Schemes
//	@Description	Update a book
//	@Tags			Books
//	@Param			id	path	int	false	"id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	BookOutput
//	@Router			/books/{id} [patch]
//	@Security		BearerAuth
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

//	@BasePath	/api/v1
//
// Books godoc
//
//	@Summary	Delete a book
//	@Schemes
//	@Description	Delete a book
//	@Tags			Books
//	@Param			id	path	int	false	"id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	BookOutput
//	@Router			/books/{id} [delete]
//	@Security		BearerAuth
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
