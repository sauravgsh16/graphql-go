package service

import (
	"github.com/sauravgsh16/graphql-go/domain"
	"github.com/sauravgsh16/graphql-go/domain/author"
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

// Create receiver for Authorservice
// @TODO: Complete
func (as *AuthorService) Create(domain.Object) (domain.Object, error) {
	return nil, nil
}

// Delete receiver for Authorservice
// @TODO: Complete
func (as *AuthorService) Delete(id int) error {
	return nil
}
