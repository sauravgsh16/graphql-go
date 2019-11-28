package resolver

import (
	"github.com/graphql-go/graphql"
)

// GraphQLResolver interface
type GraphQLResolver interface {
	IndividualResolver(p graphql.ResolveParams) (interface{}, error)
	AllResolver(p graphql.ResolveParams) (interface{}, error)
}

// AllResolvers returns a list of GraphQLResolver
func AllResolvers() []GraphQLResolver {
	return []GraphQLResolver{&AuthorResolver{}, &BookResolver{}}
}
