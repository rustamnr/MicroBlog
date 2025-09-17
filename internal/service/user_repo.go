package service

import (
	"sync"

	customErrors "github.com/lsmltesting/MicroBlog/internal/errors"
	"github.com/lsmltesting/MicroBlog/internal/models"
)

type UserRepository interface {
	Save(user *models.User) (int, error)
	FindUserById(id int) (*models.User, error)
}

type inMemoryUserRepo struct {
	mtx    sync.RWMutex
	data   map[int]*models.User
	lastID int
}

func NewInMemoryUserRepo() UserRepository {
	return &inMemoryUserRepo{
		data: make(map[int]*models.User),
	}
}

func (r *inMemoryUserRepo) Save(user *models.User) (int, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	r.lastID++
	user.ID = r.lastID
	r.data[r.lastID] = user

	return r.lastID, nil
}

func (r *inMemoryUserRepo) FindUserById(id int) (*models.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	user, ok := r.data[id]
	if !ok {
		return nil, customErrors.ErrNotFindUser
	}

	return user, nil
}
