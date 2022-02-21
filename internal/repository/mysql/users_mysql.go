package mysql

import (
	"database/sql"
	"fmt"
	"go_web_server/internal/model"
)

type UserDB struct {
	db *sql.DB
}

func NewUserPoll(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

func (u UserDB) AddUserRepo(user model.User) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (user_name,age,city) VALUES (?,?,?);`, Users)

	res, err := u.db.Exec(query, user.Name, user.Age, user.City)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (u UserDB) FindAllUsersRepo() ([]model.User, error) {
	var userList []model.User

	query := fmt.Sprintf(`SELECT * FROM %s;`, Users)

	res, err := u.db.Query(query)
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

func (u UserDB) AddBookRepo(userID, bookID int) error {
	query := fmt.Sprintf(`INSERT INTO %s (user_id,book_id) VALUES (?,?);`, ReadBooks)

	err, _ := u.db.Query(query, userID, bookID)
	if err != nil {
		return err.Err()
	}

	return nil
}

func (u UserDB) DeleteBookFromUserRepo(userID, bookID int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE  user_id = ? AND book_id = ?;`, ReadBooks)
	if err, _ := u.db.Query(query, userID, bookID); err != nil {
		return err.Err()
	}

	return nil
}

func (u UserDB) UpdateUserRepo(user model.User) (int, error) {
	query := fmt.Sprintf(`UPDATE %s SET user_name = ?, age = ?,city = ? WHERE id=?;`, Users)

	res, err := u.db.Exec(query, user.Name, user.Age, user.City, user.ID)
	if err != nil {
		return 0, err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowAffected), nil
}

func (u UserDB) DeleteUserRepo(id int) error {
	deleteUserWithBooks := fmt.Sprintf(`DELETE FROM %s WHERE  user_id = %d`, ReadBooks, id)
	query := fmt.Sprintf(`DELETE FROM %s where id = ?;`, Users)

	_, err := u.db.Exec(deleteUserWithBooks)
	if err != nil {
		return err
	}

	_, err = u.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (u UserDB) FindUserByIDRepo(id int) (model.User, error) {
	var user model.User

	query := fmt.Sprintf(`SELECT * from %s where id = ?`, Users)

	res, _ := u.db.Query(query, id)
	for res.Next() {
		if err := res.Scan(&user.ID, &user.Name, &user.Age, &user.City); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}
