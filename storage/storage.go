package storage

import (
	"errors"
	"time"

	"github.com/xakimjonov/article/modules"
)

var InMemoryArticleData []modules.Article

func CreateArticle(id string, entity modules.MakeArticle) error {
	var article modules.Article

	article.Id = id
	article.Content = entity.Content
	article.AuthorId = entity.AuthorId
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)
	return nil

}

func GetArticleById(id string) (modules.Article, error) {
	for _, v := range InMemoryArticleData {
		if v.Id == id {
			return v, nil
		}
	}
	return modules.Article{}, errors.New("article not found")
}

func GetArticleList() (resp []modules.Article, err error) {
	resp = InMemoryArticleData
	return resp, err
}

func DeleteArticle(id string) (modules.Article, error) {
	for i, v := range InMemoryArticleData {
		if v.Id == id {
			InMemoryArticleData = remove(InMemoryArticleData, i)
			return v, nil
		}
	}
	return modules.Article{}, errors.New("article not found")
}

func remove(slice []modules.Article, s int) []modules.Article {
	return append(slice[:s], slice[s+1:]...)
}

func UpadateArticle(article modules.UpadateArticle) error{
	for i, v := range InMemoryArticleData {
		if v.Id == article.Id  {
            v.Content = article.Content
			t := time.Now()
			v.UpdatedAt = &t
			InMemoryArticleData[i] = v
			return nil
}
	}
	return errors.New("article not found")
}