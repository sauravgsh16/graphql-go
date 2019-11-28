package schema

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/sauravgsh16/graphql-go/resolver"
)

// Schema GraphQL
var Schema graphql.Schema

// MustInit initializes the GraphQL schema, panics otherwise
func MustInit(r ...resolver.GraphQLResolver) {
	var queryFields graphql.Fields

	queryFields = make(map[string]*graphql.Field, 0)

	// Add query fields
	for _, rv := range r {
		switch rv.(type) {
		case *resolver.AuthorResolver:
			for name, field := range getAuthors(rv) {
				queryFields[name] = field
			}

		case *resolver.BookResolver:
			for name, field := range getBooks(rv) {
				queryFields[name] = field
			}
		default:
			panic(fmt.Sprintf("Resolver %T: Not Implemented", rv))
		}
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: queryFields,
	}

	mutation := graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: getAuthorMutationFields(),
	}
	// @TODO: Need to add mutation for books

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(mutation),
	}

	var err error
	Schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(fmt.Errorf("failed to initialize schema: %v", err))
	}
	log.Println("GraphQL schema initialized")
}
