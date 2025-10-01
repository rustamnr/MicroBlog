package user

import (
	"sync"

	customErrors "github.com/lsmltesting/MicroBlog/internal/errors"
	"github.com/lsmltesting/MicroBlog/internal/models"
)

type UserRepository interface {
	Save(user *models.User) (int, error)
	FindUserByID(ID int) (*models.User, error)
	UpdatePostHistory(userID int, postID int) error
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

func (r *inMemoryUserRepo) FindUserByID(ID int) (*models.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	user, ok := r.data[ID]
	if !ok {
		return nil, customErrors.ErrNotFindUser
	}

	return user, nil
}

func (r *inMemoryUserRepo) UpdatePostHistory(userID int, postID int) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	user, ok := r.data[userID]
	if !ok {
		return customErrors.ErrNotFindUser
	}

	if user.PostHistory == nil {
		user.PostHistory = make(map[int]struct{})
	}

	user.PostHistory[postID] = struct{}{}

	return nil
}
