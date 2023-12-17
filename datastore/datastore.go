package datastore

import (
	"database/sql"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"library/model"
)

type lib struct{}

func New() *lib {
	return &lib{}
}

func (s *lib) GetByID(ctx *gofr.Context, id int) (*model.Book, error) {
	var resp model.Book
	strId := strconv.Itoa(id)
	err := ctx.DB().QueryRowContext(ctx, " SELECT BookID, Title, AuthorID, Issued FROM Books where BookID=?", strId).
		Scan(&resp.BookID, &resp.Title, &resp.AuthorID, &resp.Issued)
	switch err {
	case sql.ErrNoRows:
		strId := strconv.Itoa(id)
		return &model.Book{}, errors.EntityNotFound{Entity: "Books", ID: strId}
	case nil:
		return &resp, nil
	default:
		return &model.Book{}, err
	}
}

func (s *lib) Create(ctx *gofr.Context, library *model.Book) (*model.Book, error) {
	var resp model.Book

	res, err := ctx.DB().ExecContext(ctx, "INSERT INTO Books (Title, AuthorID, Issued) VALUES (?,?,?)", library.Title, library.AuthorID, library.Issued)

	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}
	err = ctx.DB().QueryRowContext(ctx, "SELECT * FROM Books WHERE BookID = ?", id).Scan(&resp.BookID, &resp.Title, &resp.AuthorID, &resp.Issued)
	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}

	return &resp, nil
}

func (s *lib) Update(ctx *gofr.Context, library *model.Book) (*model.Book, error) {

	_, err := ctx.DB().ExecContext(ctx, "UPDATE Books SET Title=?, AuthorID=?, Issued=? WHERE BookID=?", library.Title, library.AuthorID, library.Issued, library.BookID)
	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}

	return library, nil
}

func (s *lib) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM Books where BookID=?", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
