package service

import (
	"fmt"

	"github.com/sauravgsh16/graphql-go/domain"
	"github.com/sauravgsh16/graphql-go/domain/author"
	"github.com/sauravgsh16/graphql-go/domain/book"
)

// AuthServ to interact with service
var AuthServ Service = &AuthorService{}

// AuthorService struct
type AuthorService struct{}

// Get receiver for Authorservice
func (as *AuthorService) Get(id int) (domain.Object, error) {
	a := &author.Author{}
	if err := a.Select(id); err != nil {
		return nil, err
	}
	return a, nil
}

// GetAll returns all the authors
func (as *AuthorService) GetAll() ([]domain.Object, error) {
	dao := &author.Author{}
	authors, err := dao.SelectAll()
	if err != nil {
		return nil, err
	}

	var objs []domain.Object
	for _, a := range authors {
		objs = append(objs, a)
	}

	return objs, nil
}

// GetAllByKey returns list of contents by key
func (as *AuthorService) GetAllByKey(auid int) (interface{}, error) {
	doa := &book.Book{}
	books, err := doa.SelectAllByAuthorID(auid)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// Create receiver for Authorservice
// @TODO: Complete
func (as *AuthorService) Create(d domain.Object) (domain.Object, error) {
	a, ok := d.(*author.Author)
	if !ok {
		return nil, fmt.Errorf("Invalid (%T) in create author", d)
	}
	if err := a.Insert(); err != nil {
		return nil, err
	}
	return a, nil
}

// Delete receiver for Authorservice
// @TODO: Complete
func (as *AuthorService) Delete(id int) error {
	a, err := as.Get(id)
	if err != nil {
		return err
	}
	if err := a.(*author.Author).Delete(id); err != nil {
		return err
	}
	return nil
}
