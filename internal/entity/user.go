package entity

import (
	"errors"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          int
	DisplayName string
	Email       string
	Password    string
	Image       string
}

const defaultImage = "https://i.imgur.com/FXmcoOQ.png"

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)

func NewUser(displayName, email, password string) (*User, error) {
	encPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed hashing password: ", err)
	}

	u := &User{
		DisplayName: displayName,
		Email:       email,
		Password:    string(encPwd),
		Image:       defaultImage,
	}

	if err := u.ValidateUserFields(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) ValidateUserFields() error {
	if u.DisplayName == "" {
		return errors.New("display name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if !emailRegex.MatchString(u.Email) {
		return errors.New("email is invalid")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
