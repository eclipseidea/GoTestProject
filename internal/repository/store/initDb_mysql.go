package store

import (
	"database/sql"
	"fmt"
)

var CreateTableUserQuery = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
                                id        INT AUTO_INCREMENT PRIMARY KEY,
                                user_name  VARCHAR(40),
                                age       INT,
                                city      VARCHAR(255));`, Users)

var CreateTableBookQuery = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
                                id        INT AUTO_INCREMENT PRIMARY KEY,
                                book_name VARCHAR(40),
                                genre     VARCHAR(255),
                                author    VARCHAR(255));`, Books)

var CreateTableReadBooks = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
                                user_id INT NOT NULL,
                                book_id INT NOT NULL,
	                            FOREIGN KEY (user_id) REFERENCES users (id) 
                                     ON DELETE RESTRICT ON UPDATE CASCADE,
	                            FOREIGN KEY (book_id) REFERENCES books (id) 
                                     ON DELETE RESTRICT ON UPDATE CASCADE)`, ReadBooks)

var dropTablesQuery = fmt.Sprintf(`DROP TABLE IF EXISTS %s,%s,%s`, ReadBooks, Books, Users)

var insertIntoTableBooks = fmt.Sprintf(`INSERT INTO %s(book_name,genre,author) 
                                       VALUES ('Super book','Fantasy','Tom'),
                                              ('Good book','Comedy','Steve'),
                                              ('Bad book','Horror','Roman'),
                                              ('Grait book','Medicine','Semen'),
                                              ('Programs','Programming','Vasya')`, Books)

var insertIntoTableUsers = fmt.Sprintf(`INSERT INTO %s(user_name,age,city) 
                                       VALUES ('Nikola',44,'Lutsk'),
                                              ('Roman',41,'Lvov'),
                                              ('Vovan',23,'Kiev')`, Users)

var insertIntoTableReadBooks = fmt.Sprintf(`INSERT INTO %s(user_id,book_id) 
                                       VALUES (3,3),(3,1),(2,1),(1,5)`, ReadBooks)

type InitDBRepository struct {
	db *sql.DB
}

func InitTables(db *sql.DB) *InitDBRepository {
	return &InitDBRepository{db: db}
}

func (i InitDBRepository) DropDataBase() error {
	_, err := i.db.Exec(dropTablesQuery)
	if err != nil {
		return err
	}

	return nil
}

func (i InitDBRepository) CreateTables() error {
	err := i.DropDataBase()
	if err != nil {
		return err
	}

	_, err = i.db.Exec(CreateTableBookQuery)
	if err != nil {
		return err
	}

	_, err = i.db.Exec(CreateTableUserQuery)
	if err != nil {
		return err
	}

	_, err = i.db.Exec(CreateTableReadBooks)
	if err != nil {
		return err
	}

	return nil
}

func (i InitDBRepository) InsertInto() error {
	_, err := i.db.Exec(insertIntoTableBooks)
	if err != nil {
		return err
	}

	_, err = i.db.Exec(insertIntoTableUsers)
	if err != nil {
		return err
	}

	_, err = i.db.Exec(insertIntoTableReadBooks)
	if err != nil {
		return err
	}

	return nil
}
