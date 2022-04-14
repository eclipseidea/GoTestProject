package store

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_NewRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := NewRepository(db)

	assert.NotNil(t, db)
	assert.NotEmpty(t, db)
	assert.NotNil(t, repo)
	assert.NotEmpty(t, repo)
	assert.Nil(t, mock.ExpectationsWereMet())
}
