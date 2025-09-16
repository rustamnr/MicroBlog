package models

import (
	"net/mail"
	"regexp"

	customErrors "github.com/lsmltesting/MicroBlog/internal/errors"
)

type User struct {
	userName string `json:"username"`
	email    string `json:"emal"`
	id       int    `json:"int"`
}

func NewUser(username string, email string, id int) *User {
	return &User{
		userName: username,
		email:    email,
		id:       id,
	}
}

func (user *User) UserName() string {
	return user.userName
}

func (user *User) Email() string {
	return user.email
}

func (user *User) Id() int {
	return user.id
}

func (user *User) SetUserName(username string) error {
	re := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9]*$`)
	if re.MatchString(username) {
		user.userName = username
		return nil
	}
	return customErrors.ErrWrongUserName
}

func (user *User) SetUserEmail(email string) error {
	addr, err := mail.ParseAddress(email)

	if err != nil {
		return customErrors.ErrWrongEmaiAddress
	}

	user.email = addr.Address
	return nil
}
