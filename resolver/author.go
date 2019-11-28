package resolver

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/sauravgsh16/graphql-go/domain/author"
	"github.com/sauravgsh16/graphql-go/service"
)

// AuthorResolver struct
type AuthorResolver struct{}

// IndividualResolver resolves individual get requests
func (r AuthorResolver) IndividualResolver(p graphql.ResolveParams) (interface{}, error) {
	author, err := service.AuthServ.Get(p.Args["id"].(int))
	if err != nil {
		return nil, err
	}
	return author, nil
}

// AllResolver resolves all get request
func (r AuthorResolver) AllResolver(p graphql.ResolveParams) (interface{}, error) {
	authors, err := service.AuthServ.GetAll()
	if err != nil {
		return nil, err
	}

	return authors, nil
}

// JoinResolver resolves the join request
func (r AuthorResolver) JoinResolver(p graphql.ResolveParams) (interface{}, error) {
	a, ok := p.Source.(*author.Author)
	if !ok {
		return nil, fmt.Errorf("Invalid type (%T) sent to execute query", p.Source)
	}
	author, err := service.AuthServ.GetAllByKey(a.ID)
	if err != nil {
		return nil, err
	}
	return author, nil
}
