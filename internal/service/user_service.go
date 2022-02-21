package service

import (
	"go_web_server/internal/model"
	"go_web_server/internal/repository"
)

type UserRepo struct {
	repo repository.UserDataRepository
}

func NewUserService(repo repository.UserDataRepository) *UserRepo {
	return &UserRepo{repo: repo}
}

func (u UserRepo) AddUser(user model.User) (int, error) {
	return u.repo.AddUserRepo(user)
}

func (u UserRepo) FindAllUsers() ([]model.User, error) {
	return u.repo.FindAllUsersRepo()
}

func (u UserRepo) AddBookToUser(userID, bookID int) error {
	return u.repo.AddBookRepo(userID, bookID)
}

func (u UserRepo) DeleteBookFromUser(userID, bookID int) error {
	return u.repo.DeleteBookFromUserRepo(userID, bookID)
}

func (u UserRepo) UpdateUser(user model.User) (int, error) {
	return u.repo.UpdateUserRepo(user)
}

func (u UserRepo) DeleteUser(id int) error {
	return u.repo.DeleteUserRepo(id)
}

func (u UserRepo) FindUserByID(id int) (model.User, error) {
	return u.repo.FindUserByIDRepo(id)
}
