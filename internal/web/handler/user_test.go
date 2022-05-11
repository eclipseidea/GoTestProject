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

func TestAddUser(t *testing.T) {
	const query = "INSERT INTO users"

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
		"POST",
		UserAddRoute,
		strings.NewReader(`{"Name":"John","Age":24,"City":"Toronto"}`))

	mock.ExpectExec(query).
		WithArgs("John", 24, "Toronto").
		WillReturnResult(sqlmock.NewResult(1, 1))

	router.POST(UserAddRoute, handler.AddUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestAddUserValidationRequestBody_Error(t *testing.T) {
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
		UserAddRoute,
		strings.NewReader(`{"Name":"John","Age":"24","City":"Toronto"}`))

	router.POST(UserAddRoute, handler.AddUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestAddUserQuery_Error(t *testing.T) {
	const query = "INSERT INTO users"

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
		UserAddRoute,
		strings.NewReader(`{"Name":"John","Age":24,"City":"Toronto"}`))

	mock.ExpectExec(query).WillReturnError(errors.New(""))

	router.POST(UserAddRoute, handler.AddUser)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 500, w.Code)
}

func TestFindAllUsers(t *testing.T) {
	const query = "SELECT * FROM users;"

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	rows := sqlmock.NewRows([]string{"Id", "Name", "Age", "City"}).
		AddRow(test.UserMock.ID, test.UserMock.Name, test.UserMock.Age, test.UserMock.City)

	mock.ExpectQuery(query).WillReturnRows(rows)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		UsersFindAllRoute,
		strings.NewReader(`nil`))

	router.GET(UsersFindAllRoute, handler.FindAllUsers)

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

func TestFindAllUsersQuery_Error(t *testing.T) {
	const query = "SELECT * FROM users;"

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	rows := sqlmock.NewRows([]string{"Id", "Name", "Age", "City"})

	mock.ExpectQuery(query).WillReturnRows(rows)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		UsersFindAllRoute,
		strings.NewReader(`nil`))

	router.GET(UsersFindAllRoute, handler.FindAllUsers)

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

func TestAddBookToUser(t *testing.T) {
	const query = "INSERT INTO read_books"

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
		"POST",
		UserAddBookRoute,
		strings.NewReader(`{ "UserId": 3, "BookId": 3}`))

	mock.ExpectExec(query).
		WithArgs(3, 3).
		WillReturnResult(sqlmock.NewResult(1, 1))

	router.POST(UserAddBookRoute, handler.AddBookToUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestAddBookToUserValidationRequestBody_Error(t *testing.T) {
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
		UserAddBookRoute,
		strings.NewReader(`{ "UserId": "3", "BookId": 3}`))

	router.POST(UserAddBookRoute, handler.AddBookToUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestAddBookToUserQueryError(t *testing.T) {
	const query = "INSERT INTO read_books"

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
		UserAddBookRoute,
		strings.NewReader(`{ "UserId": 3, "BookId": 3}`))

	mock.ExpectExec(query).WillReturnError(errors.New(""))

	router.POST(UserAddBookRoute, handler.AddBookToUser)

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

func TestDeleteBookFromUser(t *testing.T) {
	const (
		query = "DELETE FROM read_books"
		url   = "/user/db_query/deleteBookFromUser/userId/3/bookId/3/"
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
		strings.NewReader(`{3,3}`))

	mock.ExpectExec(query).
		WithArgs(3, 3).
		WillReturnResult(sqlmock.NewResult(0, 1))

	router.DELETE(UserDeleteBookRoute, handler.DeleteBookFromUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestDeleteBookFromUserQueryParamValidationUserId_Error(t *testing.T) {
	const url = "/user/db_query/deleteBookFromUser/userId/test/bookId/3/"

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
		"DELETE",
		url,
		strings.NewReader(`{"ID": 3}`))

	router.DELETE(UserDeleteBookRoute, handler.DeleteBookFromUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestDeleteBookFromUserQueryParamValidationBookId_Error(t *testing.T) {
	const url = "/user/db_query/deleteBookFromUser/userId/3/bookId/test/"

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
		"DELETE",
		url,
		strings.NewReader(`{3,3}`))

	router.DELETE(UserDeleteBookRoute, handler.DeleteBookFromUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestDeleteBookFromUserQuery_Error(t *testing.T) {
	const (
		query = "DELETE FROM read_books"
		url   = "/user/db_query/deleteBookFromUser/userId/3/bookId/3/"
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
		strings.NewReader(`{3,3}`))

	mock.ExpectExec(query).WillReturnError(errors.New(""))

	router.DELETE(UserDeleteBookRoute, handler.DeleteBookFromUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 500, w.Code)
}

func TestUpdateUser(t *testing.T) {
	const query = "UPDATE users SET user_name = ?, age = ?,city = ? WHERE id=?;"

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
		UserUpdateRoute,
		strings.NewReader(`{"Id": 2,"Name" : "Sara","Age" : 22, "City" : "Toronto"}`))

	mock.ExpectExec(query).
		WithArgs("Sara", 22, "Toronto", 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	router.PUT(UserUpdateRoute, handler.UpdateUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestUpdateUserValidationRequestBodyId_Error(t *testing.T) {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
		UserUpdateRoute,
		strings.NewReader(`{"Id": "2","Name" : "Sara","Age" : 22, "City" : "Toronto"}`))

	router.PUT(UserUpdateRoute, handler.UpdateUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestUpdateUserValidationRequestBodyAge_Error(t *testing.T) {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
		UserUpdateRoute,
		strings.NewReader(`{"Id": 2,"Name" : "Sara","Age" : "22", "City" : "Toronto"}`))

	router.PUT(UserUpdateRoute, handler.UpdateUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestUpdateUserQuery_Error(t *testing.T) {
	const query = "UPDATE users SET user_name = ?, age = ?,city = ? WHERE id=?;"

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
		UserUpdateRoute,
		strings.NewReader(`{"Id": 2,"Name" : "Sara","Age" : 22, "City" : "Toronto"}`))

	mock.ExpectExec(query).WillReturnError(errors.New(""))

	router.PUT(UserUpdateRoute, handler.UpdateUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 500, w.Code)
}

func TestDeleteUser(t *testing.T) {
	const (
		deleteUserWithBooks = "DELETE FROM read_books WHERE  user_id = ?"
		query               = "DELETE FROM users where id = ?"
		url                 = "/user/db_query/deleteUser/1/"
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
		strings.NewReader(`{"ID": 1}`))

	mock.ExpectExec(deleteUserWithBooks).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))

	router.DELETE(UserDeleteRoute, handler.DeleteUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestDeleteUserRequestBodyValidation_Error(t *testing.T) {
	const url = "/user/db_query/deleteUser/test/"

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
		"DELETE",
		url,
		strings.NewReader(`{"ID": test}`))

	router.DELETE(UserDeleteRoute, handler.DeleteUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestDeleteUserQueryRun_Error(t *testing.T) {
	const (
		deleteUserWithBooks = "DELETE FROM read_books WHERE  user_id = ?"
		query               = "DELETE FROM users where id = ?"
		url                 = "/user/db_query/deleteUser/1/"
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
		strings.NewReader(`{"ID": 1}`))

	mock.ExpectExec(deleteUserWithBooks).WithArgs("id").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(query).WithArgs("id").WillReturnResult(sqlmock.NewResult(0, 0))

	router.DELETE(UserDeleteRoute, handler.DeleteUser)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 500, w.Code)
}

func TestFindUserByID(t *testing.T) {
	const (
		query = "SELECT * from users where id = ?"
		url   = "/user/db_query/findUserById/1/"
	)

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

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	row := sqlmock.NewRows([]string{"Id", "Name", "Age", "City"}).
		AddRow(test.UserMock.ID, test.UserMock.Name, test.UserMock.Age, test.UserMock.City)

	mock.ExpectQuery(query).WillReturnRows(row)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		url,
		strings.NewReader(`nil`))

	router.GET(UserFindByIDRoute, handler.FindUserByID)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestFindUserByIDRequestBodyValidation_Error(t *testing.T) {
	const (
		url = "/user/db_query/findUserById/test/"
	)

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		url,
		strings.NewReader(`nil`))

	router.GET(UserFindByIDRoute, handler.FindUserByID)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 400, w.Code)
}

func TestFindUserByIDQueryRun_Error(t *testing.T) {
	const (
		query = "SELECT * from users where id = ?"
		url   = "/user/db_query/findUserById/1/"
	)

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

	row := sqlmock.NewRows([]string{"Id", "Name", "Age", "City"})

	mock.ExpectQuery(query).WillReturnRows(row)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		url,
		strings.NewReader(`nil`))

	router.GET(UserFindByIDRoute, handler.FindUserByID)

	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 500, w.Code)
}
