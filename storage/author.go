package storage

import (
	"time"

	"github.com/xakimjonov/article/modules"
)

var InMemoryAuthorData []modules.Author

func CreateAuthor(id string, entity modules.MakeAuthor) error {
	var author modules.Author
    
	author.Id = id
	author.Firstname = entity.Firstname
	author.Lastname = entity.Lastname
	author.CreatedAt = time.Now()

	InMemoryAuthorData = append(InMemoryAuthorData, author)
	return nil

}
