package models

import "time"

type Like struct {
	CreatedAt time.Time `json:"createdAt"`
	UserFrom  *User     `json:"userFrom"`
}

func NewLike(userFrom *User) *Like {
	return &Like{
		CreatedAt: time.Now(),
		UserFrom:  userFrom,
	}
}

// func (like *Like) CreatedAt() time.Time {
// 	return like.createdAt
// }

// func (like *Like) UserFrom() *User {
// 	return like.userFrom
// }
