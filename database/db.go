package database

import (
	"context"
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
	dbName = "graphql"
)

// DB postgres
var DB *sql.DB

// Init initializes db
func init() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPwd, dbName,
	)
	once.Do(func() {
		var err error
		DB, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		if err = DB.Ping(); err != nil {
			panic(err)
		}
	})
}

// GetConn returns a DB Connection
func GetConn(ctx context.Context) *sql.Conn {
	conn, err := DB.Conn(ctx)
	if err != nil {
		panic(err)
	}
	return conn
}

// MustExec function executes a given query. Panics on failure
func MustExec(query string, args ...interface{}) {
	_, err := DB.Exec(query, args...)
	if err != nil {
		panic(err)
	}
}
