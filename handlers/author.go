package handlers

import (
	"net/http"

	"github.com/xakimjonov/article/modules"
	"github.com/xakimjonov/article/storage"

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
func CreateAuthor(c *gin.Context) {
	var body modules.MakeAuthor
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New().String()

	err := storage.CreateAuthor(id, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author, err := storage.GetAuthorById(id)
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
// @Success     200 {object} modules.JSONResponse{data=[]modules.Author}
// @Failure     400 {object} modules.JSONErrorResponse
// @Router      /v1/author  [get]
func GetAuthorList(c *gin.Context) {
	list, err := storage.GetAuthorList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,modules.JSONResponse{
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
func GetAuthorById(c *gin.Context) {
	idStr := c.Param("id")

	author, err := storage.GetAuthorById(idStr)
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
// @Description Change current author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       article body     modules.UpadateAuthor true "author body"
// @Success     200     {object} modules.JSONResponse{data=modules.Author}
// @Failure     400     {object} modules.JSONErrorResponse
// @Router      /v1/author [put]

func UpdateAuthor(c *gin.Context) {
	var body modules.UpdateAuthor
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := storage.UpdateAuthor(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	author, err := storage.GetAuthorById(body.Id)
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
func DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")
	author, err := storage.DeleteAuthor(idStr)

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
