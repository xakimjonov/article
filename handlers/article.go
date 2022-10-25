package handlers

import (
	"net/http"
	"strconv"

	"github.com/xakimjonov/article/modules"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateArticle godoc
// @Summary     Create article
// @Description Create a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     modules.MakeArticle true "article"
// @Success     201     {object} modules.JSONResponse{data=modules.Article}
// @Failure     400     {object} modules.JSONErrorResponse
// @Router      /v1/article [post]
func (h *Handler) CreateArticle(c *gin.Context) {
	var body modules.MakeArticle
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New().String()

	err := h.Stg.CreateArticle(id, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.GetArticleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, modules.JSONResponse{
		Message: "Done sucessfully",
		Data:    article,
	})
}

// GetArticleList godoc
// @Summary     Get Article List
// @Description Get article list
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset query    int false "0"
// @Param       limit  query    int false "5"
// @Param       search query    string false "JUST"
// @Success     201 {object} modules.JSONResponse{data=[]modules.Article}
// @Failure     400 {object} modules.JSONErrorResponse
// @Router      /v1/article [get]
func (h *Handler) GetArticleList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	list, err := h.Stg.GetArticleList(offset, limit, searchStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,
		modules.JSONResponse{
			Message: "Done",
			Data:    list,
		},
	)
}

// GetArticleById godoc
// @Summary     GetArticleById
// @Description Get an article by Id
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article ID"
// @Success     200 {object} modules.JSONResponse{data=modules.ArticleFullInfo}
// @Failure     400 {object} modules.JSONErrorResponse
// @Router      /v1/article/{id} [get]
func (h *Handler) GetArticleById(c *gin.Context) {
	idStr := c.Param("id")

	article, err := h.Stg.GetArticleById(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, modules.JSONResponse{
		Message: "Done sucessfully",
		Data:    article,
	})
	return
}

// UpdatedArticle godoc
// @Summary     Update an article
// @Description Change current article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     modules.UpadateArticle true "article body"
// @Success     200     {object} modules.JSONResponse{data=modules.Article}
// @Failure     400     {object} modules.JSONErrorResponse
// @Router      /v1/article [put]
func (h *Handler) UpdatedArticle(c *gin.Context) {

	var body modules.UpadateArticle
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Stg.UpadateArticle(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.GetArticleById(body.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, modules.JSONResponse{
		Message: "Updated sucessfully",
		Data:    article,
	})
}

// DeleteArticle godoc
// @Summary     Delete Article
// @Description remove an article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article ID"
// @Success     200 {object} modules.JSONResponse{data=modules.Article}
// @Failure     400 {object} modules.JSONErrorResponse
// @Router      /v1/article/{id} [delete]
func (h *Handler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")

	article, err := h.Stg.DeleteArticle(idStr)

	if err != nil {
		c.JSON(http.StatusNotFound, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, modules.JSONResponse{
		Message: "Deleted sucessfully",
		Data:    article,
	})
	return
}
