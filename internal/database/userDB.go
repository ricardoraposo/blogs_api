package database

import (
	"database/sql"
	"gochi/internal/entity"
)

type UserDB interface {
	Create(user *entity.User) error
	GetAll() ([]*entity.User, error)
	GetByID(ID int) (*entity.User, error)
}

type userDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) UserDB {
	return &userDB{db}
}

func (u *userDB) Create(user *entity.User) error {
	stmt, err := u.db.Prepare("INSERT INTO users (display_name, email, password, image) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.DisplayName, user.Email, user.Password, user.Image)
	if err != nil {
		return err
	}

	return nil
}

func (u *userDB) GetAll() ([]*entity.User, error) {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []*entity.User
	for rows.Next() {
		var u entity.User
		err := rows.Scan(&u.ID, &u.DisplayName, &u.Email, &u.Password, &u.Image)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

func (u *userDB) GetByID(ID int) (*entity.User, error) {
	stmt, err := u.db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user entity.User
	err = stmt.QueryRow().Scan(&user.ID, &user.DisplayName, &user.Email, &user.Password, &user.Image)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
