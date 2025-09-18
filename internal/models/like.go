package models

import "time"

type Like struct {
	CreatedAt time.Time
	UserFrom  *User
}

func NewLike(userFrom *User) *Like {
	return &Like{
		CreatedAt: time.Now(),
		UserFrom:  userFrom,
	}
}
