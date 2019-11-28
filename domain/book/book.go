package book

import (
	"github.com/sauravgsh16/graphql-go/domain"
)

const (
	selectQuery       = `SELECT ID, NAME, GENRE, AUTHOR_ID FROM books WHERE ID=($1);`
	insertQuery       = `INSERT INTO books (NAME, GENRE, AUTHOR_ID ) VALUES ($1, $2, $3) RETURNING ID;`
	updateQuery       = `UPDATE books SET NAME=($1), GENRE=($2), AUTHOR_ID=($3) WHERE ID=($4);`
	deleteQuery       = `DELETE FROM books WHERE ID=($1);`
	selectAllQuery    = `SELECT ID, NAME, GENRE, AUTHOR_ID FROM books;`
	selectAuthorQuery = `SELECT ID, NAME, GENRE, AUTHOR_ID FROM books WHERE AUTHOR_ID=($1);`
)

// Book struct
type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	AuthorID int    `json:"author_id"`
}

// CanQuery returns true
func (b *Book) CanQuery() bool {
	return true
}

// Select a book from the db with matching ID
func (b *Book) Select(id int) error {
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

	if err := row.Scan(&b.ID, &b.Name, &b.Genre, &b.AuthorID); err != nil {
		return err
	}

	return nil
}

// SelectAllByAuthorID returns all books with author id
func (b *Book) SelectAllByAuthorID(authID int) ([]*Book, error) {
	conn, ctx, err := domain.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, selectAuthorQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, authID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*Book

	for rows.Next() {
		bk := &Book{}
		if err := rows.Scan(&bk.ID, &bk.Name, &bk.Genre, &bk.AuthorID); err != nil {
			return nil, err
		}
		books = append(books, bk)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

// Insert an Book into the database
func (b *Book) Insert() error {
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

	if err := stmt.QueryRowContext(ctx, b.Name, b.Genre, b.AuthorID).Scan(&returnedID); err != nil {
		return err
	}
	b.ID = returnedID
	return nil
}

// Delete an Book from the database
func (b *Book) Delete(id int) error {
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
func (b *Book) SelectAll() ([]*Book, error) {
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
		bk := &Book{}
		if err := rows.Scan(&bk.ID, &bk.Name, &bk.Genre, &bk.AuthorID); err != nil {
			return nil, err
		}

		books = append(books, bk)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
