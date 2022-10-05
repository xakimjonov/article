package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var InMemoryArticleData []Article

func remove(slice []Article, s int) []Article {
	return append(slice[:s], slice[s+1:]...)
}

func CreateArticle(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article.Id = uuid.New().String()
	article.CreatedAt = time.Now()
	InMemoryArticleData = append(InMemoryArticleData, article)

	c.JSON(http.StatusOK, gin.H{
		"data":    InMemoryArticleData,
		"message": "Article | Create",
	})
}

// GetArticleList godoc
// @Summary      List articles
// @Description  get articles
// @Tags         articles
// @Accept       json
// @Produce      json
// @Success      200  {array} JSONResponse{data=[]Article}
// @Router       /v1/article [get]
var GetArticleList = func(c *gin.Context) {
	c.JSON(http.StatusOK, 
		JSONResponse{
		Message: "Article | GetList",
		Data:    InMemoryArticleData,
	})
}

func GetArticleById(c *gin.Context) {
	idStr := c.Param("id")
	for _, v := range InMemoryArticleData {
		if v.Id == idStr {
			c.JSON(http.StatusOK, gin.H{
				"message": "Article | GetById",
				"data":    v,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Article | GetById | Not Found",
		"data":    nil,
	})
}

var UpdatedArticle = func(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, v := range InMemoryArticleData {
		if v.Id == article.Id {
			t := time.Now()
			article.UpdatedAt = &t
			article.CreatedAt = v.CreatedAt
			InMemoryArticleData[i] = article

			c.JSON(http.StatusOK, gin.H{
				"data":    InMemoryArticleData,
				"message": "Article | Update",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Article | Update | NOt Found",
		"data":    InMemoryArticleData,
	})
}

func DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	for i, v := range InMemoryArticleData {
		if v.Id == idStr {
			InMemoryArticleData = remove(InMemoryArticleData, i)
			c.JSON(http.StatusOK, gin.H{
				"message": "Article | Delete",
				"data":    v,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Article | Delete | Not Found",
	})
}
