package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/sauravgsh16/graphql-go/service"
)

type authorResolver struct{}

func (r *authorResolver) IndividualResolver(p graphql.ResolveParams) (interface{}, error) {
	author, err := service.AuthServ.Get(p.Args["id"].(int))
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (r *authorResolver) AllResolver(p graphql.ResolveParams) (interface{}, error) {
	authors, err := service.AuthServ.GetAll()
	if err != nil {
		return nil, err
	}
	return authors, nil
}
