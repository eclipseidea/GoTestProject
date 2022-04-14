package store

import (
	"errors"
	"go_web_server/internal/model"
	"go_web_server/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddUserRepo(t *testing.T) {
	const query = "INSERT INTO users"

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

	repo := NewUserPoll(db)

	mock.ExpectExec(query).
		WithArgs(test.UserMock.Name, test.UserMock.Age, test.UserMock.City).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := repo.AddUserRepo(test.UserMock)

	assert.Equal(t, 1, id)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestAddUserRepoQuery_Error(t *testing.T) {
	const query = "INSERT INTO users"

	_error := errors.New("add user query error")

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

	repo := NewUserPoll(db)

	mock.ExpectExec(query).
		WithArgs(test.UserMock.Name, test.UserMock.Age, test.UserMock.City).
		WillReturnError(_error)

	id, err := repo.AddUserRepo(test.UserMock)
	if err != nil {
		return
	}

	assert.Equal(t, 0, id)
	assert.Equal(t, _error, err)
	assert.Nil(t, err)
	assert.Error(t, err)
	assert.Nil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindAllUsersRepo(t *testing.T) {
	const query = "SELECT * FROM users;"

	var userList []model.User

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

	repo := NewUserPoll(db)

	rows := sqlmock.NewRows([]string{"Id", "Name", "Age", "City"}).
		AddRow(test.UserMock.ID, test.UserMock.Name, test.UserMock.Age, test.UserMock.City)

	mock.ExpectQuery(query).WillReturnRows(rows)

	userList, err = repo.FindAllUsersRepo()

	assert.NoError(t, err)
	assert.NotEmpty(t, userList)
	assert.Len(t, userList, 1)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindAllUsers_Error(t *testing.T) {
	const query = "SELECT * FROM users;"

	var userList []model.User

	_error := errors.New("find all users query err")

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

	repo := NewUserPoll(db)

	mock.ExpectExec(query).
		WillReturnError(_error)

	userList, err = repo.FindAllUsersRepo()
	if err != nil {
		return
	}

	assert.Error(t, err)
	assert.Equal(t, _error, err)
	assert.Empty(t, userList)
	assert.Len(t, userList, 0)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUserAddBookRepo(t *testing.T) {
	const query = "INSERT INTO read_books"

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

	repo := NewUserPoll(db)

	mock.ExpectExec(query).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UserAddBookRepo(1, 1)

	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUserAddBookRepoQuery_Error(t *testing.T) {
	const query = "INSERT INTO read_books"

	_error := errors.New("query error")

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

	repo := NewUserPoll(db)

	mock.ExpectExec(query).
		WithArgs(1, 1).
		WillReturnError(_error)

	err = repo.UserAddBookRepo(1, 1)

	assert.NotNil(t, err)
	assert.Equal(t, _error, err)
	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteBookFromUserRepo(t *testing.T) {
	const query = "DELETE FROM read_books"

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

	repo := NewUserPoll(db)

	mock.ExpectExec(query).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 0))

	err = repo.DeleteBookFromUserRepo(1, 1)

	assert.Equal(t, err, nil)
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteBookFromUserRepoQuery_Error(t *testing.T) {
	const query = "DELETE FROM read_books"

	_error := errors.New("query error")

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

	repo := NewUserPoll(db)

	mock.ExpectExec(query).
		WithArgs(1, 1).
		WillReturnError(_error)

	err = repo.DeleteBookFromUserRepo(1, 1)

	assert.Equal(t, err, _error)
	assert.Error(t, err)
	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateUserRepo(t *testing.T) {
	const query = "UPDATE users SET user_name = ?, age = ?,city = ? WHERE id=?;"

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

	repo := NewUserPoll(db)

	mock.ExpectExec(query).WithArgs(test.UserMock.Name, test.UserMock.Age, test.UserMock.City, test.UserMock.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	row, err := repo.UpdateUserRepo(test.UserMock)
	if err != nil {
		return
	}

	assert.Equal(t, 1, row)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateUserRepo_QueryError(t *testing.T) {
	const query = "UPDATE users SET;"

	_error := errors.New("query error")

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

	repo := NewUserPoll(db)

	mock.ExpectExec(query).WithArgs().
		WillReturnError(_error)

	row, err := repo.UpdateUserRepo(test.UserMock)
	if err != nil {
		return
	}

	assert.Equal(t, _error, err)
	assert.Equal(t, 0, row)
	assert.Error(t, err)
	assert.NotNil(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteUserRepo(t *testing.T) {
	const (
		deleteUserWithBooks = "DELETE FROM read_books"
		query               = "DELETE FROM users"
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

	repo := NewUserPoll(db)

	mock.ExpectExec(deleteUserWithBooks).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.DeleteUserRepo(1)
	if err != nil {
		return
	}

	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteUserRepoQuery_Error(t *testing.T) {
	const deleteUserWithBooks = "DELETE FROM read_books"

	_error := errors.New("query error")

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

	repo := NewUserPoll(db)

	mock.ExpectExec(deleteUserWithBooks).WillReturnError(_error)

	err = repo.DeleteUserRepo(1)
	if err != nil {
		return
	}

	assert.NotNil(t, repo)
	assert.Error(t, err)
	assert.Equal(t, _error, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteUserRepoSecondQuery_Error(t *testing.T) {
	const (
		deleteUserWithBooks = "DELETE FROM read_books"
		query               = "DELETE FROM users"
	)

	_error := errors.New("query error")

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

	repo := NewUserPoll(db)

	mock.ExpectExec(deleteUserWithBooks).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(query).WillReturnError(_error)

	err = repo.DeleteUserRepo(1)
	if err != nil {
		return
	}

	assert.NotNil(t, repo)
	assert.Error(t, err)
	assert.Equal(t, _error, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindUserByIDRepo(t *testing.T) {
	const query = "SELECT * from users where id = ?"

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

	repo := NewUserPoll(db)

	row := sqlmock.NewRows([]string{"Id", "Name", "Age", "City"}).
		AddRow(test.UserMock.ID, test.UserMock.Name, test.UserMock.Age, test.UserMock.City)

	mock.ExpectQuery(query).WithArgs(test.UserMock.ID).WillReturnRows(row)

	user, err := repo.FindUserByIDRepo(1)

	assert.NoError(t, err)
	assert.NotNil(t, db)
	assert.NotEmpty(t, user)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindUserByIDRepoQuery_Error(t *testing.T) {
	const query = "SELECT * from users where id = ?"

	_error := errors.New("query error")

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

	repo := NewUserPoll(db)

	row := sqlmock.NewRows([]string{"Id", "Name", "Age", "City"})

	mock.ExpectQuery(query).WithArgs(test.UserMock.ID).
		WillReturnRows(row).
		WillReturnError(_error)

	user, err := repo.FindUserByIDRepo(1)

	assert.Error(t, err)
	assert.Equal(t, _error, err)
	assert.NotNil(t, db)
	assert.Empty(t, user)
	assert.Nil(t, mock.ExpectationsWereMet())
}
