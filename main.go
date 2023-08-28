package main

import (
	auth "github.com/fahmiyonda007/go-gin-gorm/controllers/auth"
	controllers "github.com/fahmiyonda007/go-gin-gorm/controllers/books"
	"github.com/fahmiyonda007/go-gin-gorm/middleware"
	"github.com/fahmiyonda007/go-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.POST("/api/login", auth.Login)

	book := r.Group("/api/books")
	book.Use(middleware.EnsureValidToken())
	{
		book.GET("", controllers.Books)
		book.GET("/:id", controllers.Book)
		book.POST("/:id", controllers.CreateBook)
		book.PATCH("/:id", controllers.UpdateBook)
		book.DELETE("/:id", controllers.DeleteBook)
	}

	r.Run()
}
