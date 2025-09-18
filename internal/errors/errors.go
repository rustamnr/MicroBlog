package errors

import "errors"

var (
	ErrWrongUserName     = errors.New("wrong user name")
	ErrWrongEmailAddress = errors.New("wrong email address")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrNotFindUser       = errors.New("not find user")

	ErrEmptyPostText          = errors.New("empty text for post")
	ErrPostLikeAlreadyCreated = errors.New("like is already created")
	ErrNotFindPost            = errors.New("not find post")
)
