package handlers

import (
	"net/http"
	"strconv"

	"github.com/xakimjonov/article/modules"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAuthorgodoc
// @Summary     Create author
// @Description Create a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     modules.MakeAuthor true "author"
// @Success     201    {object} modules.JSONResponse{data=modules.Author}
// @Failure     400    {object} modules.JSONErrorResponse
// @Router      /v1/author [post]
func (h *Handler) CreateAuthor(c *gin.Context) {
	var body modules.MakeAuthor
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New().String()

	err := h.IM.CreateAuthor(id, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author, err := h.IM.GetAuthorById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, modules.JSONResponse{
		Message: "Done sucessfully",
		Data:    author,
	})
}

// GetauthorList godoc
// @Summary     Get AuthorList
// @Description Get Authorlist
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       offset query    int false "0"
// @Param       limit  query    int false "5"
// @Param       search query    string false "JUST"
// @Success     200 {object} modules.JSONResponse{data=[]modules.Author}
// @Failure     400 {object} modules.JSONErrorResponse
// @Router      /v1/author  [get]
func (h *Handler) GetAuthorList(c *gin.Context) {
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

	list, err := h.IM.GetAuthorList(offset, limit, searchStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, modules.JSONResponse{
		Message: "Done",
		Data:    list,
	})
}

// GetauthorById godoc
// @Summary     GetauthorById
// @Description Get an Author by Id
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       id  path     string true "AuthorID"
// @Success     200 {object} modules.JSONResponse{data=modules.ArticleFullInfo}
// @Failure     400 {object} modules.JSONErrorResponse
// @Router      /v1/author/{id} [get]
func (h *Handler) GetAuthorById(c *gin.Context) {
	idStr := c.Param("id")

	author, err := h.IM.GetAuthorById(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, modules.JSONResponse{
		Message: "Done sucessfully",
		Data:    author,
	})
	return
}

// UpdateAuthor godoc
// @Summary     Update an author
// @Description Update an author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       article body     modules.UpdateAuthor true "author body"
// @Success     200     {object} modules.JSONResponse{data=modules.Author}
// @Failure     400     {object} modules.JSONErrorResponse
// @Router      /v1/author [put]
func (h *Handler) UpdateAuthor(c *gin.Context) {
	var body modules.UpdateAuthor
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.IM.UpdateAuthor(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	author, err := h.IM.GetAuthorById(body.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, modules.JSONResponse{
		Message: "Updated sucessfully",
		Data:    author,
	})
}

// DeleteAuthor godoc
// @Summary     Delete author
// @Description remove an author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       id  path     string true "AuthorID"
// @Success     200 {object} modules.JSONResponse{data=modules.Author}
// @Failure     400 {object} modules.JSONErrorResponse
// @Router      /v1/author/{id} [delete]
func (h *Handler) DeleteAuthor(c *gin.Context) {

	idStr := c.Param("id")
	author, err := h.IM.DeleteAuthor(idStr)

	if err != nil {
		c.JSON(http.StatusNotFound, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, modules.JSONResponse{
		Message: "Deleted sucessfully",
		Data:    author,
	})
	return
}
