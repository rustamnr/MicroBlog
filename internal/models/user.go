package models

import (
	"net/mail"
	"regexp"
	"unicode"

	customErrors "github.com/lsmltesting/MicroBlog/internal/errors"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ID       int    `json:"id"`
}

func NewUser(username string, email string, password string) (*User, error) {
	user := &User{
		Username: username,
		Email:    email,
		Password: password,
	}

	if err := user.SetUsername(username); err != nil {
		return nil, customErrors.ErrWrongUserName
	}

	if err := user.SetUserEmail(email); err != nil {
		return nil, customErrors.ErrWrongEmailAddress
	}

	if err := user.SetPassword(password); err != nil {
		return nil, customErrors.ErrInvalidPassword
	}

	return user, nil
}

func (user *User) SetUsername(username string) error {
	re := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9]*$`)
	if re.MatchString(username) {
		user.Username = username
		return nil
	}
	return customErrors.ErrWrongUserName
}

func (user *User) SetUserEmail(email string) error {
	addr, err := mail.ParseAddress(email)

	if err != nil {
		return customErrors.ErrWrongEmailAddress
	}

	user.Email = addr.Address
	return nil
}

func (user *User) SetPassword(password string) error {
	if len(password) < 8 {
		return customErrors.ErrInvalidPassword
	}

	var (
		hasLower = false
		hasUpper = false
		hasDigit = false
	)

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if !hasLower || !hasUpper || !hasDigit {
		return customErrors.ErrInvalidPassword
	}

	user.Password = password
	return nil
}
