package author

import (
	"github.com/sauravgsh16/graphql-go/domain"
)

const (
	selectQuery    = `SELECT id, name, age FROM authors WHERE ID=($1);`
	insertQuery    = `INSERT INTO authors (NAME, AGE) VALUES ($1, $2) RETURNING ID;`
	updateQuery    = `UPDATE authors SET NAME=($1), AGE=($2) WHERE ID=($3);`
	deleteQuery    = `DELETE FROM authors WHERE ID=($1);`
	selectAllQuery = `SELECT ID, NAME, AGE FROM authors;`
)

// Author struct
type Author struct {
	ID   int
	Name string
	Age  int
}

// CanQuery returns true
func (a *Author) CanQuery() bool {
	return true
}

// Select an author from the db with matching ID
func (a *Author) Select(id int) error {
	conn, ctx, err := domain.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, selectQuery)
	if err != nil {
		return err
	}

	row := stmt.QueryRowContext(ctx, id)

	if err := row.Scan(&a.ID, &a.Name, &a.Age); err != nil {
		return err
	}

	return nil
}

// Insert an author into the database
func (a *Author) Insert() error {
	conn, ctx, err := domain.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, insertQuery)
	if err != nil {
		return err
	}

	var returnedID int

	if err := stmt.QueryRowContext(ctx, a.Name, a.Age).Scan(&returnedID); err != nil {
		return err
	}
	a.ID = returnedID
	return nil
}

// Delete an author from the database
func (a *Author) Delete(id int) error {
	conn, ctx, err := domain.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, deleteQuery)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}
	return nil
}

// SelectAll returns all authors from db
func (a *Author) SelectAll() ([]*Author, error) {
	conn, ctx, err := domain.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, selectAllQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*Author

	for rows.Next() {
		au := &Author{}
		if err := rows.Scan(&au.ID, &au.Name, &au.Age); err != nil {
			return nil, err
		}

		authors = append(authors, au)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}
