package models

import "time"

type Like struct {
	createdAt time.Time `json: "createdAt"`
	userFrom  *User     `json: "userFrom"`
}

func NewLike(userFrom *User) *Like {
	return &Like{
		createdAt: time.Now(),
		userFrom:  userFrom,
	}
}

func (like *Like) CreatedAt() time.Time {
	return like.createdAt
}

func (like *Like) UserFrom() *User {
	return like.userFrom
}
