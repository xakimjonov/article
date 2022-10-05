package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/xakimjonov/article/docs"
)

// @contact.name   API Article
// @contact.url    https://t.me/xakimjonov_0008
// @contact.email  xakimjonov0008@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	docs.SwaggerInfo.Title = "Swagger Article API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"

	InMemoryArticleData = make([]Article, 0)
	router := gin.Default()

	v1 := router.Group("/v1/article")
	{
		v1.POST("/", CreateArticle)
		v1.GET("/", GetArticleList)
		v1.GET("/:id", GetArticleById)
		v1.PUT("/", UpdatedArticle)
		v1.DELETE("/:id", DeleteArticle)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":5454") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
