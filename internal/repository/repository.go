package repository

import (
	"database/sql"
	"go_web_server/internal/model"
	"go_web_server/internal/repository/mysql"
)

type InitDBRepository interface {
	CreateTables() error
	InsertInto() error
	DropDataBase() error
}

type UserDataRepository interface {
	AddUserRepo(user model.User) (int, error)
	FindAllUsersRepo() ([]model.User, error)
	AddBookRepo(userID, bookID int) error
	DeleteBookFromUserRepo(userID, bookID int) error
	UpdateUserRepo(user model.User) (int, error)
	DeleteUserRepo(id int) error
	FindUserByIDRepo(id int) (model.User, error)
}

type BookDataRepository interface {
	AddBookRepo(book model.Book) (int, error)
	UpdateBookRepo(book model.Book) (int, error)
	DeleteBookRepo(id int) error
	FindAllBooksRepo() ([]model.Book, error)
	FindBookByNameRepo(name string) (model.Book, error)
}

type Repository struct {
	InitDBRepository
	UserDataRepository
	BookDataRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		InitDBRepository:   mysql.InitTables(db),
		UserDataRepository: mysql.NewUserPoll(db),
		BookDataRepository: mysql.NewBookPool(db),
	}
}
