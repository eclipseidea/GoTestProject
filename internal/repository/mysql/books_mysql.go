package mysql

import (
	"database/sql"
	"fmt"
	"go_web_server/internal/model"
)

type BookDB struct {
	db *sql.DB
}

func NewBookPool(db *sql.DB) *BookDB {
	return &BookDB{db: db}
}

func (b BookDB) AddBookRepo(book model.Book) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (book_name,genre,author) VALUES (?,?,?);`, Books)
	res, err := b.db.Exec(query, book.Name, book.Genre, book.Author)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (b BookDB) UpdateBookRepo(book model.Book) (int, error) {
	query := fmt.Sprintf(`UPDATE %s SET book_name = ?, genre = ?,author = ? WHERE id=?;`, Books)
	res, err := b.db.Exec(query, book.Name, book.Genre, book.Author, book.ID)

	if err != nil {
		return 0, err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowAffected), nil
}

func (b BookDB) DeleteBookRepo(id int) error {
	setForeignKeyChecksFalse := `SET FOREIGN_KEY_CHECKS = OFF`
	setForeignKeyChecksTrue := `SET FOREIGN_KEY_CHECKS= ON`
	query := fmt.Sprintf(`DELETE FROM %s where id = ?;`, Books)

	_, err := b.db.Exec(setForeignKeyChecksFalse)
	if err != nil {
		return err
	}

	_, err = b.db.Exec(query, id)
	if err != nil {
		return err
	}

	_, err = b.db.Exec(setForeignKeyChecksTrue)
	if err != nil {
		return err
	}

	return nil
}

func (b BookDB) FindAllBooksRepo() ([]model.Book, error) {
	var bookList []model.Book

	query := fmt.Sprintf(`SELECT * FROM %s;`, Books)

	res, err := b.db.Query(query)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var book model.Book

		err := res.Scan(&book.ID, &book.Name, &book.Genre, &book.Author)
		if err != nil {
			return nil, err
		}

		bookList = append(bookList, book)
	}

	return bookList, nil
}

func (b BookDB) FindBookByNameRepo(name string) (model.Book, error) {
	var book model.Book

	query := fmt.Sprintf(`SELECT * from %s WHERE book_name = ?;`, Books)

	res, _ := b.db.Query(query, name)
	for res.Next() {
		if err := res.Scan(&book.ID, &book.Name, &book.Genre, &book.Author); err != nil {
			return model.Book{}, err
		}
	}

	return book, nil
}
