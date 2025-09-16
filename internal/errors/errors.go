package errors

import "errors"

var (
	ErrWrongUserName    = errors.New("wrong user name")
	ErrWrongEmaiAddress = errors.New("wrong email address")
)
