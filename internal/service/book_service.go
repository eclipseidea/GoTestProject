package service

import (
	"go_web_server/internal/model"
	"go_web_server/internal/repository"
)

type BookRepo struct {
	repo repository.BookDataRepository
}

func NewBookService(repo repository.BookDataRepository) *BookRepo {
	return &BookRepo{repo: repo}
}

func (b BookRepo) AddBook(book model.Book) (int, error) {
	return b.repo.AddBookRepo(book)
}

func (b BookRepo) UpdateBook(book model.Book) (int, error) {
	return b.repo.UpdateBookRepo(book)
}

func (b BookRepo) DeleteBook(id int) error {
	return b.repo.DeleteBookRepo(id)
}

func (b BookRepo) FindAllBooks() ([]model.Book, error) {
	return b.repo.FindAllBooksRepo()
}

func (b BookRepo) FindBookByName(name string) (model.Book, error) {
	return b.repo.FindBookByNameRepo(name)
}
