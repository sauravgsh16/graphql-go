package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/graphql-go/graphql"

	"github.com/sauravgsh16/graphql-go/resolver"
	"github.com/sauravgsh16/graphql-go/schema"
)

func GraphQLHanlder() gin.HandlerFunc {
	r := resolver.NewResolver()
	schema.MustInit(r)
}

/*

Use map of GraphQLResolver

With type switching to associate correct resolver

*/
