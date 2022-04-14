package store

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDropDataBase(t *testing.T) {
	const query = "DROP TABLE IF EXISTS read_books,books,users"

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

	repo := InitTables(db)

	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.DropDataBase()
	if err == nil {
		return
	}

	assert.Equal(t, nil, err)
	assert.NotEmpty(t, repo)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDropDataBase_Error(t *testing.T) {
	const query = "DROP TABLE IF EXISTS"

	_error := errors.New("drop db query error")

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

	repo := InitTables(db)

	mock.ExpectExec(query).WillReturnError(_error)

	err = repo.DropDataBase()
	if err == nil {
		return
	}

	assert.Equal(t, _error, err)
	assert.NotEmpty(t, repo)
	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreateTables(t *testing.T) {
	const (
		query            = "DROP TABLE IF EXISTS read_books,books,users"
		createTableQuery = "CREATE TABLE IF NOT EXISTS"
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

	repo := InitTables(db)

	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(createTableQuery).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(createTableQuery).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(createTableQuery).WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.CreateTables()

	assert.NotEmpty(t, repo)
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreateTables_ErrorDropDataBase(t *testing.T) {
	const query = "DROP TABLE IF EXISTS"

	_error := errors.New("drop db error")

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

	repo := InitTables(db)

	mock.ExpectExec(query).WillReturnError(_error)

	err = repo.CreateTables()

	assert.NotEmpty(t, repo)
	assert.Equal(t, _error, err)
	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreateTables_ErrorCreateTablesUserQuery(t *testing.T) {
	const (
		query                = "DROP TABLE IF EXISTS"
		createTableUserQuery = "CREATE TABLE IF NOT EXISTS"
	)

	_error := errors.New("user table create error")

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

	repo := InitTables(db)

	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(createTableUserQuery).WillReturnError(_error)

	err = repo.CreateTables()

	assert.NotEmpty(t, repo)
	assert.Equal(t, _error, err)
	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreateTables_ErrorCreateTablesBooKQuery(t *testing.T) {
	const (
		query            = "DROP TABLE IF EXISTS"
		createTableQuery = "CREATE TABLE IF NOT EXISTS"
	)

	_error := errors.New("book table create error")

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

	repo := InitTables(db)

	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(createTableQuery).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(createTableQuery).WillReturnError(_error)

	err = repo.CreateTables()

	assert.NotEmpty(t, repo)
	assert.Equal(t, _error, err)
	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreateTables_ErrorCreateTablesReadBooKQuery(t *testing.T) {
	const (
		query            = "DROP TABLE IF EXISTS"
		createTableQuery = "CREATE TABLE IF NOT EXISTS"
	)

	_error := errors.New("book table create error")

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

	repo := InitTables(db)

	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(createTableQuery).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(createTableQuery).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(createTableQuery).WillReturnError(_error)

	err = repo.CreateTables()

	assert.NotEmpty(t, repo)
	assert.Equal(t, _error, err)
	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestInsertInto(t *testing.T) {
	const (
		insertIntoTableBooks     = "INSERT INTO books"
		insertIntoTableUsers     = "INSERT INTO users"
		insertIntoTableReadBooks = "INSERT INTO read_books"
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

	repo := InitTables(db)

	mock.ExpectExec(insertIntoTableBooks).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(insertIntoTableUsers).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(insertIntoTableReadBooks).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.InsertInto()
	if err != nil {
		return
	}

	assert.NotEmpty(t, repo)
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestInsertIntoTableBooks_Error(t *testing.T) {
	const insertIntoTableBooks = "INSERT INTO books"

	_error := errors.New("insert into table books error")

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

	repo := InitTables(db)

	mock.ExpectExec(insertIntoTableBooks).WillReturnError(_error)

	err = repo.InsertInto()
	if err != nil {
		return
	}

	assert.NotEmpty(t, repo)
	assert.Equal(t, _error, err)
	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestInsertIntoTableUsers_Error(t *testing.T) {
	const (
		insertIntoTableBooks = "INSERT INTO books"
		insertIntoTableUsers = "INSERT INTO users"
	)

	_error := errors.New("insert into table users error")

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

	repo := InitTables(db)

	mock.ExpectExec(insertIntoTableBooks).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(insertIntoTableUsers).WillReturnError(_error)

	err = repo.InsertInto()
	if err != nil {
		return
	}

	assert.NotEmpty(t, repo)
	assert.Equal(t, _error, err)
	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestInsertIntoTableReadBooks_Error(t *testing.T) {
	const (
		insertIntoTableBooks     = "INSERT INTO books"
		insertIntoTableUsers     = "INSERT INTO users"
		insertIntoTableReadBooks = "INSERT INTO read_books"
	)

	_error := errors.New("insert into table read_books error")

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

	repo := InitTables(db)

	mock.ExpectExec(insertIntoTableBooks).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(insertIntoTableUsers).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(insertIntoTableReadBooks).WillReturnError(_error)

	err = repo.InsertInto()
	if err != nil {
		return
	}

	assert.NotEmpty(t, repo)
	assert.Equal(t, _error, err)
	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}
