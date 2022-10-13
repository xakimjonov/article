package inmemeory

import (
	"errors"
	"strings"
	"time"

	"github.com/xakimjonov/article/modules"
)

func (im InMemeory) CreateArticle(id string, entity modules.MakeArticle) error {
	var article modules.Article

	article.Id = id
	article.Content = entity.Content
	article.AuthorId = entity.AuthorId
	article.CreatedAt = time.Now()

	im.Db.InMemoryArticleData = append(im.Db.InMemoryArticleData, article)
	return nil

}

func (im InMemeory) GetArticleById(id string) (modules.ArticleFullInfo, error) {
	var result modules.ArticleFullInfo
	for _, v := range im.Db.InMemoryArticleData {
		if v.Id == id {
			author, err := im.GetAuthorById(v.AuthorId)
			if err != nil {
				return result, err
			}
			result.Id = v.Id
			result.Content = v.Content
			result.Author = author
			result.CreatedAt = v.CreatedAt
			result.UpdatedAt = v.UpdatedAt
			return result, nil
		}
	}
	return modules.ArticleFullInfo{}, errors.New("article not found")
}

func (im InMemeory) GetArticleList(offset, limit int, search string) (resp []modules.Article, err error) {
	off := 0
	c := 0
	for _, v := range im.Db.InMemoryArticleData {
		if strings.Contains(v.Title, search) || strings.Contains(v.Body, search) {
			if offset <= off {
				c++
				resp = append(resp, v)
			}

			if c >= limit {
				break
			}

			off++
		}
	}

	return resp, err
}
func (in InMemeory) DeleteArticle(id string) (modules.Article, error) {
	for i, v := range in.Db.InMemoryArticleData {
		if v.Id == id {
			in.Db.InMemoryArticleData = remove(in.Db.InMemoryArticleData, i)
			return v, nil
		}
	}
	return modules.Article{}, errors.New("article not found")
}

func remove(slice []modules.Article, s int) []modules.Article {
	return append(slice[:s], slice[s+1:]...)
}

func (in InMemeory) UpadateArticle(article modules.UpadateArticle) error {
	for i, v := range in.Db.InMemoryArticleData {
		if v.Id == article.Id {
			v.Content = article.Content
			t := time.Now()
			v.UpdatedAt = &t
			in.Db.InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("article not found")
}
