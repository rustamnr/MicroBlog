package post

import (
	"sync"

	customErrors "github.com/lsmltesting/MicroBlog/internal/errors"
	"github.com/lsmltesting/MicroBlog/internal/models"
)

type PostRepository interface {
	Save(post *models.Post) (int, error)
	FindPostByID(postID int) (*models.Post, error)
	GetAllPosts() (map[int]*models.Post, error)
	UpdateLikeHistory(postID int, likeID int) error
}

type inMemoryPostRepo struct {
	mtx    sync.RWMutex
	data   map[int]*models.Post
	lastID int
}

func NewInMemoryPostRepo() PostRepository {
	return &inMemoryPostRepo{
		data: make(map[int]*models.Post),
	}
}

func (r *inMemoryPostRepo) Save(post *models.Post) (int, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	r.lastID++

	post.ID = r.lastID
	r.data[r.lastID] = post

	return r.lastID, nil
}

func (r *inMemoryPostRepo) FindPostByID(postID int) (*models.Post, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	post, ok := r.data[postID]
	if !ok {
		return nil, customErrors.ErrEmptyPostText
	}
	return post, nil
}

func (r *inMemoryPostRepo) GetAllPosts() (map[int]*models.Post, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if len(r.data) == 0 {
		return nil, customErrors.ErrNotAnyPostExists
	}
	return r.data, nil
}

func (r *inMemoryPostRepo) UpdateLikeHistory(postID int, likeID int) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	_, ok := r.data[postID]
	if !ok {
		return customErrors.ErrNotFindPost
	}

	r.data[postID].HistoryLikes[likeID] = struct{}{}

	return nil
}
