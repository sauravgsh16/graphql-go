package database

import (
	"database/sql"
	"fmt"
	"sync"

	// postgres driver
	_ "github.com/lib/pq"
)

var once sync.Once

const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPwd  = "postgres"
	dbName = "library"
)

// DB postgres
var DB *sql.DB

// Init initializes db
func Init() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPwd, dbName,
	)
	once.Do(func() {
		DB, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		if err := DB.Ping(); err != nil {
			panic(err)
		}
	})
}

// MustExec function executes a given query. Panics on failure
func MustExec(query string, args ...interface{}) {
	_, err := DB.Exec(query, args...)
	if err != nil {
		panic(err)
	}
}
