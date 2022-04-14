package store

import (
	"database/sql"
	"fmt"
	"go_web_server/internal/model"
	"time"

	"golang.org/x/net/context"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserPoll(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (s *UserRepository) AddUserRepo(user model.User) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (user_name,age,city) VALUES (?,?,?);`, Users)

	res, err := s.db.Exec(query, user.Name, user.Age, user.City)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *UserRepository) FindAllUsersRepo() ([]model.User, error) {
	var userList []model.User

	query := fmt.Sprintf(`SELECT * FROM %s;`, Users)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var user model.User

		err := res.Scan(&user.ID, &user.Name, &user.Age, &user.City)
		if err != nil {
			return nil, err
		}

		userList = append(userList, user)
	}

	return userList, nil
}

func (s *UserRepository) UserAddBookRepo(userID, bookID int) error {
	query := fmt.Sprintf(`INSERT INTO %s (user_id,book_id) VALUES (?,?);`, ReadBooks)

	_, err := s.db.Exec(query, userID, bookID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserRepository) DeleteBookFromUserRepo(userID, bookID int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE  user_id = ? AND book_id = ?;`, ReadBooks)
	if _, err := s.db.Exec(query, userID, bookID); err != nil {
		return err
	}

	return nil
}

func (s *UserRepository) UpdateUserRepo(user model.User) (int, error) {
	query := fmt.Sprintf(`UPDATE %s SET user_name = ?, age = ?,city = ? WHERE id=?;`, Users)

	res, err := s.db.Exec(query, user.Name, user.Age, user.City, user.ID)
	if err != nil {
		return 0, err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowAffected), nil
}

func (s *UserRepository) DeleteUserRepo(id int) error {
	deleteUserWithBooks := fmt.Sprintf(`DELETE FROM %s WHERE  user_id = %d`, ReadBooks, id)
	query := fmt.Sprintf(`DELETE FROM %s where id = ?;`, Users)

	_, err := s.db.Exec(deleteUserWithBooks)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserRepository) FindUserByIDRepo(id int) (model.User, error) {
	var user model.User

	query := fmt.Sprintf(`SELECT * from %s where id = ?`, Users)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Age, &user.City)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
