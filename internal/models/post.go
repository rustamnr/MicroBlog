package models

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
	createdAt time.Time `json: "createdAt"`
	userFrom  *User     `json: "userFrom"`
}

type Post struct {
	createdAt time.Time     `json: "createdAt"`
	updatedAt time.Time     `json: "updatedAt"`
	text      string        `json: "test"`
	user      *User         `json: "user"`
	likes     map[int]*Like `json: "likes"`
	id        uuid.UUID     `json: id`
}

func NewPost(user *User, text string) *Post {
	return &Post{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		text:      text,
		user:      user,
		likes:     make(map[int]*Like),
		id:        uuid.New(),
	}
}

func (post *Post) CreatedAt() time.Time {
	return post.createdAt
}

func (post *Post) UpdateAt() time.Time {
	return post.updatedAt
}

func (post *Post) Text() string {
	return post.text
}

func (post *Post) User() *User {
	return post.user
}

func (post *Post) Likes() map[int]*Like {
	return post.likes
}

func (post *Post) Id() uuid.UUID {
	return post.id
}
