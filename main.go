package main

import (
	"log"
	"os"

	auth "github.com/fahmiyonda007/go-gin-gorm/controllers/auth"
	controllers "github.com/fahmiyonda007/go-gin-gorm/controllers/books"
	docs "github.com/fahmiyonda007/go-gin-gorm/docs"
	"github.com/fahmiyonda007/go-gin-gorm/middleware"
	"github.com/fahmiyonda007/go-gin-gorm/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Gin Book Service
//	@version		1.0
//	@description	A book management service API in Go using Gin framework GORM and Auth0.
//	@termsOfService	https://tos.santoshk.dev

//	@contact.name	Fahmi Yonda
//	@contact.url	https://instagram.com/fahmiyonda
//	@contact.email	yondadf04@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1
//	@schemes	http

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	r := gin.Default()

	models.ConnectDatabase()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", auth.Login)

		book := v1.Group("/books")
		book.Use(middleware.EnsureValidToken())
		{
			book.GET("", controllers.Books)
			book.GET("/:id", controllers.Book)
			book.POST("/:id", controllers.CreateBook)
			book.PATCH("/:id", controllers.UpdateBook)
			book.DELETE("/:id", controllers.DeleteBook)
		}
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)))

	r.Run(":" + os.Getenv("APP_PORT"))
}
