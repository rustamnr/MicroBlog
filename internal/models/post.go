package models

import (
	"time"

	"github.com/lsmltesting/MicroBlog/internal/errors"
)

type Post struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Text      string
	User      *User
	Likes     map[int]*Like
	ID        int
}

func NewPost(user *User, text string) (*Post, error) {
	post := &Post{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User:      user,
		Likes:     make(map[int]*Like),
	}

	// Set text for post after validating
	if err := post.SetText(text); err != nil {
		return nil, err
	}

	return post, nil
}

func (post *Post) SetText(text string) error {
	if text == "" {
		return errors.ErrEmptyPostText
	}

	post.Text = text
	post.UpdatedAt = time.Now()
	return nil
}

func (post *Post) SetLike(userId int, like *Like) error {
	// Check if like from userId is already created
	if _, ok := post.Likes[userId]; !ok {
		return errors.ErrPostLikeAlreadyCreated
	}

	post.Likes[userId] = like
	return nil
}
