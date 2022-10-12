package storage

import (
	"errors"
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

func GetAuthorById(id string) (modules.Author, error) {
	var result modules.Author
	for _, v := range InMemoryAuthorData {
		if v.Id == id {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

func GetAuthorList() (resp []modules.Author, err error) {
	resp = InMemoryAuthorData
	return resp, err
}

func UpdateAuthor(author modules.UpdateAuthor) error {
	for i, v := range InMemoryAuthorData {
		if v.Id == author.Id {
			v.Firstname = author.Firstname
			v.Lastname = author.Lastname
			t := time.Now()
			v.UpdatedAt = &t
			InMemoryAuthorData[i] = v
			return nil
		}
	}
	return errors.New("article not found")
}

func DeleteAuthor(id string) (modules.Author, error) {
	for i, v := range InMemoryAuthorData {
		if v.Id == id {
			InMemoryAuthorData = del(InMemoryAuthorData, i)
			return v, nil
		}
	}
	return modules.Author{}, errors.New("article not found")
}

func del(slice []modules.Author, s int) []modules.Author {
	return append(slice[:s], slice[s+1:]...)
}
