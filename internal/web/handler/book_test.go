package handler

import (
	"errors"
	"go_web_server/internal/repository/store"
	"go_web_server/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	const query = "INSERT INTO books"

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"POST",
		BookAddRoute,
		strings.NewReader(`{"Name": "Mafia","Genre": "fantasy","Author":"Roman"}`))

	mock.ExpectExec(query).
		WithArgs(test.BookMock.Name, test.BookMock.Genre, test.BookMock.Author).
		WillReturnResult(sqlmock.NewResult(1, 1))

	router.POST(BookAddRoute, handler.AddBook)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	assert.Equal(t, 200, w.Code)
}

func TestAddBookRequestBodyValidation_Error(t *testing.T) {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"POST",
		BookAddRoute,
		strings.NewReader(`{"Name": 1,"Genre": "fantasy","Author":"Roman"}`))

	router.POST(BookAddRoute, handler.AddBook)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestAddBookQueryRun_Error(t *testing.T) {
	const query = "INSERT INTO books"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	req, _ := http.NewRequest(
		"POST",
		BookAddRoute,
		strings.NewReader(`{"Name": "Mafia","Genre": "fantasy","Author":"Roman"}`))

	w := httptest.NewRecorder()

	mock.ExpectExec(query).WillReturnError(errors.New(""))

	router.POST(BookAddRoute, handler.AddBook)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 500, w.Code)
}

func TestUpdateBook(t *testing.T) {
	const query = "UPDATE books SET book_name = ?, genre = ?,author = ? WHERE id=?;"

	router := gin.Default()
	gin.SetMode(gin.TestMode)

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

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"PUT",
		BookUpdateRoute,
		strings.NewReader(`{"Name": "Mafia","Genre":  "fantasy","Author": "Roman","ID":1}`))

	mock.ExpectExec(query).
		WithArgs(test.BookMock.Name, test.BookMock.Genre, test.BookMock.Author, test.BookMock.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	router.PUT(BookUpdateRoute, handler.UpdateBook)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestUpdateBookRequestBodyValidation_Error(t *testing.T) {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"PUT",
		BookUpdateRoute,
		strings.NewReader(`{"Name": "Mafia","Genre":  "fantasy","Author": "Roman","ID":"1"}`))

	router.PUT(BookUpdateRoute, handler.UpdateBook)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestUpdateBookQueryRun_Error(t *testing.T) {
	const query = "UPDATE books SET"

	router := gin.Default()
	gin.SetMode(gin.TestMode)

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

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	req, _ := http.NewRequest(
		"PUT",
		BookUpdateRoute,
		strings.NewReader(`{"Name": "Mafia","Genre":  "fantasy","Author": "Roman","ID":1}`))

	w := httptest.NewRecorder()

	mock.ExpectExec(query).WillReturnError(errors.New(""))

	router.PUT(BookUpdateRoute, handler.UpdateBook)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 500, w.Code)
}

func TestDeleteBook(t *testing.T) {
	const (
		setForeignKeyChecksFalse = "SET FOREIGN_KEY_CHECKS = OFF"
		setForeignKeyChecksTrue  = "SET FOREIGN_KEY_CHECKS = ON"
		query                    = "DELETE FROM books where id = ?"
		url                      = "/book/db_query/deleteBook/5/"
	)

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	w := httptest.NewRecorder()

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	req, _ := http.NewRequest(
		"DELETE",
		url,
		strings.NewReader(`{"ID": 5}`))

	mock.ExpectExec(setForeignKeyChecksFalse).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(setForeignKeyChecksTrue).WillReturnResult(sqlmock.NewResult(0, 0))

	router.DELETE(BookDeleteRoute, handler.DeleteBook)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestDeleteBookRequestBodyValidation_Error(t *testing.T) {
	const url = "/book/db_query/deleteBook/test/"

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	req, _ := http.NewRequest(
		"DELETE",
		url,
		strings.NewReader(`{"ID": 5}`))

	router.DELETE(BookDeleteRoute, handler.DeleteBook)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestDeleteBookQueryRun_Error(t *testing.T) {
	const (
		setForeignKeyChecksFalse = "SET FOREIGN_KEY_CHECKS = OFF"
		setForeignKeyChecksTrue  = "SET FOREIGN_KEY_CHECKS = ON"
		query                    = "DELETE FROM books where id = ?"
		url                      = "/book/db_query/deleteBook/5/"
	)

	router := gin.Default()
	gin.SetMode(gin.TestMode)

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

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"DELETE",
		url,
		strings.NewReader(`{"ID": 5}`))

	mock.ExpectExec(setForeignKeyChecksFalse).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(query).WillReturnError(errors.New(""))
	mock.ExpectExec(setForeignKeyChecksTrue).WillReturnResult(sqlmock.NewResult(0, 0))

	router.DELETE(BookDeleteRoute, handler.DeleteBook)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 500, w.Code)
}

func TestFindAllBook(t *testing.T) {
	const query = "SELECT * FROM books;"

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	rows := sqlmock.NewRows([]string{"Id", "Name", "Genre", "Author"}).
		AddRow(test.BookMock.ID, test.BookMock.Name, test.BookMock.Genre, test.BookMock.Author)

	mock.ExpectQuery(query).WillReturnRows(rows)

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		BooksFindAllRoute,
		strings.NewReader(`{nil}`))

	router.GET(BooksFindAllRoute, handler.FindAllBook)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestFindAllBookQueryRun_Error(t *testing.T) {
	const query = "SELECT * FROM books;"

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	mock.ExpectQuery(query).WillReturnError(errors.New(""))

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		BooksFindAllRoute,
		strings.NewReader(`nil`))

	router.GET(BooksFindAllRoute, handler.FindAllBook)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 500, w.Code)
}

func TestFindBookByName(t *testing.T) {
	const (
		query = "SELECT * from books WHERE book_name = ?;"
		url   = "/book/db_query/findBookByName/name/"
	)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	rows := sqlmock.NewRows([]string{"Id", "Name", "Genre", "Author"}).
		AddRow(test.BookMock.ID, test.BookMock.Name, test.BookMock.Genre, test.BookMock.Author)

	mock.ExpectQuery(query).WillReturnRows(rows)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		url,
		strings.NewReader(`nil`))

	router.GET(BookFindByNameRoute, handler.FindBookByName)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestFindBookByNameQueryRun_Error(t *testing.T) {
	const (
		query = "SELECT * from books WHERE book_name = ?;"
		url   = "/book/db_query/findBookByName/name/"
	)

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	mock.ExpectQuery(query).WillReturnError(errors.New(""))

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		url,
		strings.NewReader(`nil`))

	router.GET(BookFindByNameRoute, handler.FindBookByName)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	assert.Equal(t, 500, w.Code)
}
