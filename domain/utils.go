package domain

import (
	"context"
	"database/sql"

	"github.com/sauravgsh16/graphql-go/database"
)

// GetConn returns a connection to the database and context
func GetConn() (*sql.Conn, context.Context, error) {
	ctx := context.Background()
	conn := database.GetConn(ctx)
	return conn, ctx, nil
}
