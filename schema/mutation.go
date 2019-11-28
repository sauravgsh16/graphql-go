package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/sauravgsh16/graphql-go/domain/author"
	"github.com/sauravgsh16/graphql-go/service"
)

func getAuthorMutationFields() graphql.Fields {
	return graphql.Fields{
		"createAuthor": &graphql.Field{
			Type: authorType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "Author's name",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"age": &graphql.ArgumentConfig{
					Description: "Author's age",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				age := p.Args["age"].(int)
				a := &author.Author{
					Name: name,
					Age:  age,
				}
				// @TODO: Need to find better way to do it. Without passing 'a' to create
				return service.AuthServ.Create(a)
			},
		},
		"deleteAuthor": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Author's Id",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				if err := service.AuthServ.Delete(id); err != nil {
					return nil, err
				}
				return true, nil
			},
		},
	}
}
