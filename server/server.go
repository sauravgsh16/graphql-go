package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravgsh16/graphql-go/handler"
)

var (
	router = gin.Default()
)

// StartApp starts the application
func StartApp() {
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func mapUrls() {
	router.POST("/graphql", handler.GraphQLHanlder())
}
