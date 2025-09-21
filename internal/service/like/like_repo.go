package like

import (
	"sync"

	"github.com/lsmltesting/MicroBlog/internal/errors"
	"github.com/lsmltesting/MicroBlog/internal/models"
)

type LikeRepository interface {
	Save(like *models.Like) (int, error)
	FindLikeById(likeID int) (*models.Like, error)
}

type inMemoryLikeRepo struct {
	mtx    sync.RWMutex
	data   map[int]*models.Like
	lastID int
}

func NewInMemoryLikeRepo() LikeRepository {
	return &inMemoryLikeRepo{
		data: make(map[int]*models.Like),
	}
}

func (l *inMemoryLikeRepo) Save(like *models.Like) (int, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	l.lastID++

	like.ID = l.lastID
	l.data[l.lastID] = like

	return l.lastID, nil
}

func (l *inMemoryLikeRepo) FindLikeById(likeID int) (*models.Like, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	like, ok := l.data[likeID]
	if !ok {
		return nil, errors.ErrNotFindLike
	}

	return like, nil
}
