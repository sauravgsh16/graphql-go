package service

import (
	"github.com/sauravgsh16/graphql-go/domain"
	"github.com/sauravgsh16/graphql-go/domain/author"
	"github.com/sauravgsh16/graphql-go/domain/book"
)

// BookServ to interact with service
var BookServ Service = &BookService{}

// BookService struct
type BookService struct{}

// Get receiver for Bookservice
func (bs *BookService) Get(id int) (domain.Object, error) {
	a := &book.Book{}
	if err := a.Select(id); err != nil {
		return nil, err
	}
	return a, nil
}

// GetAll returns all the books
func (bs *BookService) GetAll() ([]domain.Object, error) {
	dao := &book.Book{}
	books, err := dao.SelectAll()
	if err != nil {
		return nil, err
	}

	var objs []domain.Object
	for _, b := range books {
		objs = append(objs, b)
	}

	return objs, nil
}

// GetAllByKey return author by author_id
func (bs *BookService) GetAllByKey(auid int) (interface{}, error) {
	a := &author.Author{}
	err := a.Select(auid)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Create receiver for Bookservice
// @TODO: Complete
func (bs *BookService) Create(domain.Object) (domain.Object, error) {
	return nil, nil
}

// Delete receiver for Bookservice
// @TODO: Complete
func (bs *BookService) Delete(id int) error {
	return nil
}
