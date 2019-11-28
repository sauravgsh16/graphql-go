package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/sauravgsh16/graphql-go/database"
	"github.com/sauravgsh16/graphql-go/server"
)

func main() {
	// Init DB tables
	for _, qry := range readSQL() {
		database.MustExec(qry)
	}
	// Start server
	server.StartApp()
}

func readSQL() []string {
	content, err := ioutil.ReadFile("./database/db.sql")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	var lines []string
	for _, l := range strings.Split(string(content), ";") {
		lines = append(lines, fmt.Sprintf("%s;", strings.TrimSpace(l)))
	}
	return lines
}
