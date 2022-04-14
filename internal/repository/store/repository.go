package store

import (
	"database/sql"
	"go_web_server/internal/model"
)

type InitDB interface {
	CreateTables() error
	InsertInto() error
	DropDataBase() error
}

type UserData interface {
	AddUserRepo(user model.User) (int, error)
	FindAllUsersRepo() ([]model.User, error)
	UserAddBookRepo(userID, bookID int) error
	DeleteBookFromUserRepo(userID, bookID int) error
	UpdateUserRepo(user model.User) (int, error)
	DeleteUserRepo(id int) error
	FindUserByIDRepo(id int) (model.User, error)
}

type BookData interface {
	AddBookRepo(book model.Book) (int, error)
	UpdateBookRepo(book model.Book) (int, error)
	DeleteBookRepo(id int) error
	FindAllBooksRepo() ([]model.Book, error)
	FindBookByNameRepo(name string) (model.Book, error)
}

type Repository struct {
	InitDB
	UserData
	BookData
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		InitDB:   InitTables(db),
		UserData: NewUserPoll(db),
		BookData: NewBookPool(db),
	}
}
