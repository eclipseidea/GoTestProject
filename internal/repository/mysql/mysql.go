package mysql

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

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewMySQLDB(c *Config) (*sql.DB, error) {
	log.Println(c.Host, c.Port, c.DBName)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
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
