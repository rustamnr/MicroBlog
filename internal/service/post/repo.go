package post

import (
	"sync"

	customErrors "github.com/lsmltesting/MicroBlog/internal/errors"
	"github.com/lsmltesting/MicroBlog/internal/models"
)

type PostRepository interface {
	Save(post *models.Post) (int, error)
	FindPostById(postId int) (*models.Post, error)
	GetAllPosts() (map[int]*models.Post, error)
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
	r.data[r.lastID] = post

	return r.lastID, nil
}

func (r *inMemoryPostRepo) FindPostById(postId int) (*models.Post, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	post, ok := r.data[postId]
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
