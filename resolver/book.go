package resolver

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/sauravgsh16/graphql-go/domain/book"
	"github.com/sauravgsh16/graphql-go/service"
)

// BookResolver struct
type BookResolver struct{}

// IndividualResolver resolves individual get requests
func (r BookResolver) IndividualResolver(p graphql.ResolveParams) (interface{}, error) {
	book, err := service.BookServ.Get(p.Args["id"].(int))
	if err != nil {
		return nil, err
	}
	return book, nil
}

// AllResolver resolves all get request
func (r BookResolver) AllResolver(p graphql.ResolveParams) (interface{}, error) {
	books, err := service.BookServ.GetAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

// JoinResolver resolves all join request
func (r BookResolver) JoinResolver(p graphql.ResolveParams) (interface{}, error) {
	b, ok := p.Source.(*book.Book)
	if !ok {
		return nil, fmt.Errorf("Invalid type (%T) sent to execute query", p.Source)
	}
	books, err := service.BookServ.GetAllByKey(b.AuthorID)
	if err != nil {
		return nil, err
	}
	return books, nil
}
