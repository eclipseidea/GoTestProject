package store

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var testConfig = &ConfigDB{
	DB:       "mysql",
	Username: "root",
	Password: "Root#1234",
	Host:     "localhost",
	Port:     "3306",
	DBName:   "testDB",
}

func TestNewDB(t *testing.T) {
	testConfig = &ConfigDB{
		DB:       "mysql",
		Username: "root",
		Password: "Root#1234",
		Host:     "localhost",
		Port:     "3306",
		DBName:   "testDB",
	}

	db, err := NewDB(testConfig)

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	assert.Nil(t, err)
	assert.Equal(t, nil, err)
	assert.NotNil(t, db)
	assert.NoError(t, err)
}

func TestNewDBOpen_Error(t *testing.T) {
	testConfig = &ConfigDB{
		Username: "username",
		Password: "password",
		Host:     "db.host",
		Port:     "db.port",
		DBName:   "db.dbname",
	}

	db, err := NewDB(testConfig)

	assert.Nil(t, db)
	assert.Error(t, err)
}

func TestNewDBPing_Error(t *testing.T) {
	testConfig = &ConfigDB{
		DB:       "mysql",
		Username: "username",
		Password: "password",
		Host:     "db.host",
		Port:     "db.port",
		DBName:   "db.dbname",
	}

	db, err := NewDB(testConfig)

	assert.Nil(t, db)
	assert.Error(t, err)
}
