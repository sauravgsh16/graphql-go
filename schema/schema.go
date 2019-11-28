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
func MustInit(r resolver.GraphQLResolver) {
	var allFields graphql.Fields

	for name, field := range getAuthors(r) {
		allFields[name] = field
	}
	for name, field := range getBooks(r) {
		allFields[name] = field
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: allFields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	var err error
	Schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(fmt.Errorf("failed to initialize schema: %v", err))
	}
	log.Println("GraphQL schema initialized")
}
