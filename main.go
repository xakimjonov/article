package main

import (
	"github.com/xakimjonov/article/handlers"
	"github.com/xakimjonov/article/storage"
	"github.com/xakimjonov/article/storage/inmemeory"

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

	router := gin.Default()


	var stg storage.StoInter
	stg = inmemeory.InMemeory{
		Db: &inmemeory.DB{},
	}

	h := handlers.Handler{
	Stg: stg,
	}

	v1 := router.Group("/v1")
	{
		v1.POST("/article", h.CreateArticle)
		v1.GET("/article", h.GetArticleList)
		v1.GET("/article/:id", h.GetArticleById)
		v1.PUT("/article", h.UpdatedArticle)
		v1.DELETE("/article/:id", h.DeleteArticle)

		v1.POST("/author", h.CreateAuthor)
		v1.GET("/author", h.GetAuthorList)
		v1.GET("/author/:id", h.GetAuthorById)
		v1.PUT("/author", h.UpdateAuthor)
		v1.DELETE("/author/:id", h.DeleteAuthor)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":5454") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
