package book

import (
	"github.com/sauravgsh16/graphql-go/domain"
)

const (
	selectQuery    = `SELECT id, name, age FROM books WHERE ID=($1);`
	insertQuery    = `INSERT INTO books (NAME, AGE) VALUES ($1, $2) RETURNING ID;`
	updateQuery    = `UPDATE books SET NAME=($1), AGE=($2) WHERE ID=($3);`
	deleteQuery    = `DELETE FROM books WHERE ID=($1);`
	selectAllQuery = `SELECT ID, NAME, GENRE FROM books;`
)

// Book struct
type Book struct {
	ID    int
	Name  string
	Genre string
}

// CanQuery returns true
func (a *Book) CanQuery() bool {
	return true
}

// Select a book from the db with matching ID
func (a *Book) Select(id int) error {
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

	if err := row.Scan(&a.ID, &a.Name, &a.Genre); err != nil {
		return err
	}

	return nil
}

// Insert an Book into the database
func (a *Book) Insert(name, genre string) error {
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

	if err := stmt.QueryRowContext(ctx, name, genre).Scan(&returnedID); err != nil {
		return err
	}
	a.ID = returnedID
	return nil
}

// Delete an Book from the database
func (a *Book) Delete(id int) error {
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

// SelectAll returns all books from db
func (a *Book) SelectAll() ([]*Book, error) {
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

	var books []*Book

	for rows.Next() {
		b := &Book{}
		if err := rows.Scan(&b.ID, &b.Name, &a.Genre); err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
