package store

import (
	"errors"
	"go_web_server/internal/model"
	"go_web_server/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddBookRepo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("INSERT INTO books").
		WithArgs(test.BookMock.Name, test.BookMock.Genre, test.BookMock.Author).
		WillReturnResult(sqlmock.NewResult(1, 1))

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repo := NewBookPool(db)

	c, err := repo.AddBookRepo(test.BookMock)

	assert.Equal(t, 1, c)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestAddBookRepo_QueryError(t *testing.T) {
	const query = "INSERT INTO books"

	_error := errors.New("query run error")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	mock.ExpectExec(query).
		WithArgs(test.BookMock.Name, test.BookMock.Genre, test.BookMock.Author).
		WillReturnError(_error)

	repo := NewBookPool(db)

	c, err := repo.AddBookRepo(test.BookMock)

	assert.Equal(t, 0, c)
	assert.Equal(t, _error, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateBookRepo(t *testing.T) {
	const query = "UPDATE books SET"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := NewBookPool(db)

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	mock.ExpectExec(query).
		WithArgs(test.BookMock.Name, test.BookMock.Genre, test.BookMock.Author, test.BookMock.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	affectedRow, err := repo.UpdateBookRepo(test.BookMock)

	assert.Equal(t, 1, affectedRow)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateBookRepo_QueryError(t *testing.T) {
	const query = "UPDATE books SET"

	_error := errors.New("test query run error")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repo := NewBookPool(db)

	mock.ExpectExec(query).
		WithArgs(test.BookMock.Name, test.BookMock.Genre, test.BookMock.Author, test.BookMock.ID).
		WillReturnError(_error)

	c, err := repo.UpdateBookRepo(test.BookMock)

	assert.Equal(t, 0, c)
	assert.Equal(t, _error, err)
	assert.Error(t, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteBookRepo(t *testing.T) {
	const (
		setForeignKeyChecksFalse = "SET FOREIGN_KEY_CHECKS = OFF"
		setForeignKeyChecksTrue  = "SET FOREIGN_KEY_CHECKS = ON"
		query                    = "DELETE FROM books where id = ?"
	)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repo := NewBookPool(db)

	mock.ExpectExec(setForeignKeyChecksFalse).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(setForeignKeyChecksTrue).WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.DeleteBookRepo(1)

	assert.Nil(t, err)
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteBookRepoSetForeignKeyChecksFalse_Error(t *testing.T) {
	const setForeignKeyChecksFalse = "SET FOREIGN_KEY_CHECKS = OFF"

	_error := errors.New("set foreign key checks false query error")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repo := NewBookPool(db)

	mock.ExpectExec(setForeignKeyChecksFalse).WillReturnError(_error)

	err = repo.DeleteBookRepo(1)
	if err != nil {
		return
	}

	assert.Error(t, err)
	assert.Equal(t, _error, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteBookRepoDeleteBook_Error(t *testing.T) {
	const setForeignKeyChecksFalse = "SET FOREIGN_KEY_CHECKS = OFF"

	const query = "DELETE FROM books"

	_error := errors.New("delete book query error")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repo := NewBookPool(db)

	mock.ExpectExec(setForeignKeyChecksFalse).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(query).WillReturnError(_error)

	err = repo.DeleteBookRepo(1)
	if err != nil {
		return
	}

	assert.Error(t, err)
	assert.Equal(t, _error, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteBookRepoSetForeignKeyChecksTrue_Error(t *testing.T) {
	const (
		setForeignKeyChecksFalse = "SET FOREIGN_KEY_CHECKS = OFF"
		query                    = "DELETE FROM books"
		setForeignKeyChecksTrue  = "SET FOREIGN_KEY_CHECKS= ON"
	)

	_error := errors.New("set foreign key checks true query error")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repo := NewBookPool(db)

	mock.ExpectExec(setForeignKeyChecksFalse).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(setForeignKeyChecksTrue).WillReturnError(_error)

	err = repo.DeleteBookRepo(1)
	if err != nil {
		return
	}

	assert.Error(t, err)
	assert.Equal(t, _error, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindAllBooksRepo(t *testing.T) {
	const query = "SELECT * FROM books;"

	var bookList []model.Book

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repo := NewBookPool(db)

	rows := sqlmock.NewRows([]string{"Id", "Name", "Genre", "Author"}).
		AddRow(test.BookMock.ID, test.BookMock.Name, test.BookMock.Genre, test.BookMock.Author)

	mock.ExpectQuery(query).WillReturnRows(rows)

	bookList, err = repo.FindAllBooksRepo()

	assert.NotEmpty(t, bookList)
	assert.NoError(t, err)
	assert.NotNil(t, bookList)
	assert.Len(t, bookList, 1)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindAllBooksRepo_QueryError(t *testing.T) {
	const query = "SELECT * FROM books;"

	var bookList []model.Book

	_error := errors.New("test query error")

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repo := NewBookPool(db)

	mock.ExpectQuery(query).WillReturnError(_error)

	bookList, err = repo.FindAllBooksRepo()

	assert.Equal(t, _error, err)
	assert.Empty(t, bookList)
	assert.Error(t, err)
	assert.Len(t, bookList, 0)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindBookByNameRepo(t *testing.T) {
	const query = "SELECT * from books WHERE book_name = ?;"

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repo := NewBookPool(db)

	row := sqlmock.NewRows([]string{"Id", "Name", "Author", "Genre"}).
		AddRow(test.BookMock.ID, test.BookMock.Name, test.BookMock.Genre, test.BookMock.Author)

	mock.ExpectQuery(query).WithArgs(test.BookMock.Name).WillReturnRows(row)

	book, err := repo.FindBookByNameRepo(test.BookMock.Name)

	assert.NotEmpty(t, book)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindBookByNameRepo_Error(t *testing.T) {
	const query = "SELECT * from books WHERE book_name = ?;"

	_error := errors.New("query run error")

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repo := NewBookPool(db)

	row := sqlmock.NewRows([]string{"Id", "Name", "Author", "Genre"})

	mock.ExpectQuery(query).
		WithArgs(test.BookMock.Name).
		WillReturnRows(row).
		WillReturnError(_error)

	book, err := repo.FindBookByNameRepo(test.BookMock.Name)

	assert.Equal(t, _error, err)
	assert.Empty(t, book)
	assert.Error(t, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}
