package storage

import "github.com/xakimjonov/article/modules"

type StoInter interface {
	CreateArticle(id string, entity modules.MakeArticle) error
	GetArticleById(id string) (modules.ArticleFullInfo, error)
	DeleteArticle(id string) (modules.Article, error)
    UpadateArticle(article modules.UpadateArticle)
	GetArticleList(offset, limit int, search string) (resp []modules.Article, err error)
    

	CreateAuthor(id string, entity modules.MakeAuthor) error
	GetAuthorById(id string) (modules.Author, error)
    GetAuthorList(offset, limit int, search string) (resp []modules.Author, err error)
	UpdateAuthor(author modules.UpdateAuthor) error 
	DeleteAuthor(id string) (modules.Author, error) 

}
