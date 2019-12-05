package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	// postgres driver
	_ "github.com/lib/pq"
)

var once sync.Once

var (
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPwd  = os.Getenv("DB_PWD")
	dbName = os.Getenv("DB_NAME")
)

// DB postgres
var DB *sql.DB

// Init initializes db
func init() {
	port, err := strconv.Atoi(dbPort)
	if err != nil {
		panic(err)
	}
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, port, dbUser, dbPwd, dbName,
	)
	once.Do(func() {
		var err error
		DB, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		if err = DB.Ping(); err != nil {
			log.Println("I am breaking here")
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
