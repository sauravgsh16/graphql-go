package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/sauravgsh16/graphql-go/resolver"
)

func getBooks(r resolver.GraphQLResolver) graphql.Fields {
	var bookType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"genre": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	bookFields := graphql.Fields{
		"Book": &graphql.Field{
			Type: graphql.Type(bookType),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: r.IndividualResolver,
		},
		"Books": &graphql.Field{
			Type:    graphql.NewList(bookType),
			Resolve: r.AllResolver,
		},
	}
	return bookFields
}
