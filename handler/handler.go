package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"

	"github.com/sauravgsh16/graphql-go/resolver"
	"github.com/sauravgsh16/graphql-go/schema"
)

// GraphQLHanlder func
func GraphQLHanlder() gin.HandlerFunc {
	r := resolver.AllResolvers()
	schema.MustInit(r...)

	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
