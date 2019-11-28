package resolver

import (
	"github.com/graphql-go/graphql"
)

// GraphQLResolver interface
type GraphQLResolver interface {
	IndividualResolver(p graphql.ResolveParams) (interface{}, error)
	AllResolver(p graphql.ResolveParams) (interface{}, error)
}

type Resolver struct {
	Author GraphQLResolver
	Book   GraphQLResolver
}

// NewResolver returns new resolver
func NewResolver() *Resolver {
	return &Resolver{
		Author: &authorResolver{},
		Book:   &bookResolver{},
	}
}
