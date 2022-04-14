package store

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	Users     = "users"
	Books     = "books"
	ReadBooks = "read_books"
)

type ConfigDB struct {
	DB       string
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewDB(c *ConfigDB) (*sql.DB, error) {
	log.Println(c.Host, c.Port, c.DBName)

	db, err := sql.Open(c.DB, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.Username, c.Password, c.Host, c.Port, c.DBName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
