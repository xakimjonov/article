package inmemeory

import "github.com/xakimjonov/article/modules"

type InMemeory struct {
	Db *DB
}

type DB struct {
	InMemoryArticleData []modules.Article
	InMemoryAuthorData  []modules.Author
}
