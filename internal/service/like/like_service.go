package like

import (
	"github.com/lsmltesting/MicroBlog/internal/models"
	"github.com/lsmltesting/MicroBlog/internal/repo/like"
	"github.com/lsmltesting/MicroBlog/internal/service/post"
	"github.com/lsmltesting/MicroBlog/internal/service/user"
)

type LikeService interface {
	CreateLike(userID int, postID int) (int, error)
	GetLikeById(likeID int) (*models.Like, error)
	GetAllLikes() (map[int]*models.Like, error)
}

type likeService struct {
	repo like.LikeRepository

	userService user.UserService
	postService post.PostService
}

func NewLikeService(
	repo like.LikeRepository,
	userService user.UserService,
	postService post.PostService,
) LikeService {
	return &likeService{
		repo:        repo,
		userService: userService,
		postService: postService,
	}
}

func (l *likeService) CreateLike(userID int, postID int) (int, error) {
	// Check if user exists
	_, err := l.userService.GetUserByID(userID)
	if err != nil {
		return 0, err
	}

	// Check if post exists
	_, err = l.postService.GetPostByID(postID)
	if err != nil {
		return 0, err
	}

	like := models.NewLike(userID, postID)

	// First save like in repo. After saving like will have actual ID
	likeID, err := l.repo.Save(like)
	if err != nil {
		return 0, err
	}

	err = l.postService.UpdateLikeHistory(postID, likeID)
	if err != nil {
		return 0, err
	}

	return likeID, nil
}

func (l *likeService) GetLikeById(likeID int) (*models.Like, error) {
	return l.repo.FindLikeById(likeID)
}

func (l *likeService) GetAllLikes() (map[int]*models.Like, error) {
	return l.repo.GetAllLikes()
}
