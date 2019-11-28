package resolver

import (
	"github.com/graphql-go/graphql"
)

// GraphQLResolver interface
type GraphQLResolver interface {
	IndividualResolver(p graphql.ResolveParams) (interface{}, error)
	AllResolver(p graphql.ResolveParams) (interface{}, error)
	JoinResolver(p graphql.ResolveParams) (interface{}, error)
}

// AllResolvers returns a list of GraphQLResolver
func AllResolvers() []GraphQLResolver {
	// need to maintain order , as there is dependency during graphql type definitions
	return []GraphQLResolver{&AuthorResolver{}, &BookResolver{}}
}
