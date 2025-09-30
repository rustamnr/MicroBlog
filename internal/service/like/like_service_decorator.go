package like

import (
	"github.com/lsmltesting/MicroBlog/internal/logger"
)

type likeServiceDecorator struct {
	likeService LikeService
	lg          *logger.Logger
}

func NewLikeServiceDecorator(likeService LikeService, lg *logger.Logger) *likeServiceDecorator {
	return &likeServiceDecorator{
		likeService: likeService,
		lg:          lg,
	}
}

func (l *likeServiceDecorator) CreateLike(userID int, postID int) (int, error) {
	likeID, err := l.likeService.CreateLike(userID, postID)

	return likeID, err
}
