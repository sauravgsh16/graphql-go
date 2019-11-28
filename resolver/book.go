package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/sauravgsh16/graphql-go/service"
)

type bookResolver struct{}

func (r *bookResolver) IndividualResolver(p graphql.ResolveParams) (interface{}, error) {
	book, err := service.BookServ.Get(p.Args["id"].(int))
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookResolver) AllResolver(p graphql.ResolveParams) (interface{}, error) {
	books, err := service.BookServ.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}
