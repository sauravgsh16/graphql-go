package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/sauravgsh16/graphql-go/resolver"
)

func getAuthors(r resolver.GraphQLResolver) graphql.Fields {
	var authorType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Age": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

	authorFields := graphql.Fields{
		"Author": &graphql.Field{
			Type: graphql.Type(authorType),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: r.IndividualResolver,
		},
		"Authors": &graphql.Field{
			Type:    graphql.NewList(authorType),
			Resolve: r.AllResolver,
		},
	}
	return authorFields
}
