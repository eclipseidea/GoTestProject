package service

import (
	"go_web_server/internal/model"
	"go_web_server/internal/repository"
)

type InitDBSerVice interface {
	CreateTables() error
	InsertInto() error
}

type UserService interface {
	AddUser(user model.User) (int, error)
	FindAllUsers() ([]model.User, error)
	AddBookToUser(userID, bookID int) error
	DeleteBookFromUser(userID, bookID int) error
	UpdateUser(user model.User) (int, error)
	DeleteUser(id int) error
	FindUserByID(id int) (model.User, error)
}

type BookService interface {
	AddBook(book model.Book) (int, error)
	UpdateBook(book model.Book) (int, error)
	DeleteBook(id int) error
	FindAllBooks() ([]model.Book, error)
	FindBookByName(name string) (model.Book, error)
}

type Service struct {
	UserService
	BookService
	InitDBSerVice
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService:   NewUserService(repo.UserDataRepository),
		BookService:   NewBookService(repo.BookDataRepository),
		InitDBSerVice: NewInitDBService(repo.InitDBRepository),
	}
}
