// controllers/authors.go

package controllers

import (
	"net/http"
	"strconv"

	auth "github.com/fahmiyonda007/go-gin-gorm/controllers/auth"
	metadata "github.com/fahmiyonda007/go-gin-gorm/controllers/pagination"
	"github.com/fahmiyonda007/go-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

var resource = "author"
var readPermission = "read:" + resource
var createPermission = "create:" + resource
var updatePermission = "update:" + resource
var deletePermission = "delete:" + resource

//	@BasePath	/api/v1
//
// Authors godoc
//
//	@Summary	get all authors
//	@Schemes
//	@Description	get all authors
//	@Tags			Authors
//	@Param			page	query	int	false	"page"
//	@Param			length	query	int	false	"length"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		AuthorOutput
//	@Failure		401	{object}	handler.JSONResult
//	@Router			/authors [get]
//	@Security		BearerAuth
func Authors(c *gin.Context) {
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
	var authors []models.Author
	models.DB.Find(&authors).Count(&count).Limit(metadata.Limit(int(length))).Offset(metadata.Offset(int(page), int(length))).Find(&authors)

	meta := metadata.CalculateMetadata(int(count), int(page), int(length))
	validate := metadata.ValidateFilter(meta, int(page), int(length))

	if validate != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": validate})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"metadata": meta,
		"data":     authors,
	})
}

//	@BasePath	/api/v1
//
// Authors/:id godoc
//
//	@Summary	find a author by id
//	@Schemes
//	@Description	find a author by id
//	@Tags			Authors
//	@Param			id	path	int	true	"id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	AuthorOutput
//	@Router			/authors/{id} [get]
//	@Security		BearerAuth
func Author(c *gin.Context) {
	token, err := auth.Authenticate(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if !auth.HasPermission(token, readPermission) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		return
	}

	var author models.Author
	if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": author})
}

//	@BasePath	/api/v1
//
// Authors godoc
//
//	@Summary	Create new author
//	@Schemes
//	@Description	Create new author
//	@Tags			Authors
//	@Param			input	body	CreateAuthorInput	true	"Input"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	AuthorOutput
//	@Router			/authors [post]
//	@Security		BearerAuth
func CreateAuthor(c *gin.Context) {
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
	var input CreateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	// Create author
	author := models.Author{Name: input.Name}
	models.DB.Create(&author)

	c.JSON(http.StatusOK, gin.H{"data": author})

}

//	@BasePath	/api/v1
//
// Authors godoc
//
//	@Summary	Update a author
//	@Schemes
//	@Description	Update a author
//	@Tags			Authors
//	@Param			id		path	int					true	"id"
//	@Param			input	body	UpdateAuthorInput	false	"Input"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	AuthorOutput
//	@Router			/authors/{id} [patch]
//	@Security		BearerAuth
func UpdateAuthor(c *gin.Context) {
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
	var author models.Author
	if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Record not found!",
		})
		return
	}

	// Validate input
	var input UpdateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	models.DB.Model(&author).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": author})
}

//	@BasePath	/api/v1
//
// Authors godoc
//
//	@Summary	Delete a author
//	@Schemes
//	@Description	Delete a author
//	@Tags			Authors
//	@Param			id	path	int	true	"id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	AuthorOutput
//	@Router			/authors/{id} [delete]
//	@Security		BearerAuth
func DeleteAuthor(c *gin.Context) {
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
	var author models.Author
	if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Record not found!",
		})
		return
	}

	models.DB.Delete(&author)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
