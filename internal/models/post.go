package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Text      string        `json:"test"`
	User      *User         `json:"user"`
	Likes     map[int]*Like `json:"likes"`
	Id        uuid.UUID     `json:"id"`
}

func NewPost(user *User, text string) *Post {
	return &Post{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      text,
		User:      user,
		Likes:     make(map[int]*Like),
		Id:        uuid.New(),
	}
}

func (post *Post) SetText(text string) {
	post.Text = text
	post.UpdatedAt = time.Now()
}

func (post *Post) SetLike(userId int, like *Like) {
	post.Likes[userId] = like
}
