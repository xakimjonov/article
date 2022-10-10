package main

import (
	"github.com/xakimjonov/article/handlers"
	"github.com/xakimjonov/article/modules"
	"github.com/xakimjonov/article/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/xakimjonov/article/docs"
)

// @contact.name  API Article Test
// @contact.url   https://t.me/xakimjonov_0008
// @contact.email xakimjonov0008@gmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	storage.InMemoryArticleData = make([]modules.Article, 0)
	router := gin.Default()

	v1 := router.Group("/v1/article")
	{
		v1.POST("/", handlers.CreateArticle)
		v1.GET("/", handlers.GetArticleList)
		v1.GET("/:id", handlers.GetArticleById)
		v1.PUT("/", handlers.UpdatedArticle)
		v1.DELETE("/:id", handlers.DeleteArticle)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":5454") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
