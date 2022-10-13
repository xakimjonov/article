package inmemeory

import (
	"errors"
	"strings"
	"time"

	"github.com/xakimjonov/article/modules"
)

func (im InMemeory) CreateAuthor(id string, entity modules.MakeAuthor) error {
	var author modules.Author

	author.Id = id
	author.Firstname = entity.Firstname
	author.Lastname = entity.Lastname
	author.CreatedAt = time.Now()

	im.Db.InMemoryAuthorData = append( im.Db.InMemoryAuthorData, author)
	return nil

}

func (in InMemeory) GetAuthorById(id string) (modules.Author, error) {
	var result modules.Author
	for _, v := range in.Db.InMemoryAuthorData {
		if v.Id == id {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

func (in InMemeory) GetAuthorList(offset, limit int, search string) (resp []modules.Author, err error) {
	off := 0
	c := 0
	for _, v := range in.Db.InMemoryAuthorData {
		if strings.Contains(v.Firstname, search) || strings.Contains(v.Lastname, search) {
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

func (in InMemeory) UpdateAuthor(author modules.UpdateAuthor) error {
	for i, v := range in.Db.InMemoryAuthorData {
		if v.Id == author.Id {
			v.Firstname = author.Firstname
			v.Lastname = author.Lastname
			t := time.Now()
			v.UpdatedAt = &t
			in.Db.InMemoryAuthorData[i] = v
			return nil
		}
	}
	return errors.New("article not found")
}

func (in InMemeory) DeleteAuthor(id string) (modules.Author, error) {
	for i, v := range in.Db.InMemoryAuthorData {
		if v.Id == id {
			in.Db.InMemoryAuthorData = del(in.Db.InMemoryAuthorData, i)
			return v, nil
		}
	}
	return modules.Author{}, errors.New("article not found")
}

func del(slice []modules.Author, s int) []modules.Author {
	return append(slice[:s], slice[s+1:]...)
}
