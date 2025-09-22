package post

import (
	"github.com/lsmltesting/MicroBlog/internal/models"
	"github.com/lsmltesting/MicroBlog/internal/service/user"
)

type PostService interface {
	CreatePost(user int, text string) (int, error)
	GetPostByID(postID int) (*models.Post, error)
	GetAllPosts() (map[int]*models.Post, error)
	UpdateLikeHistory(postID int, likeID int) error
}

type postService struct {
	repo        PostRepository
	userService user.UserService
}

func NewPostService(repo PostRepository, userService user.UserService) PostService {
	return &postService{
		repo:        repo,
		userService: userService,
	}
}

func (s *postService) CreatePost(userID int, text string) (int, error) {
	// Check if user with shared userId is exists
	_, err := s.userService.GetUserByID(userID)
	if err != nil {
		return 0, err
	}

	post, err := models.NewPost(userID, text)
	if err != nil {
		return 0, err
	}

	// After creating post update user's posthistory map
	postID, err := s.repo.Save(post)
	if err != nil {
		return 0, err
	}
	err = s.userService.UpdatePostHistory(userID, postID)
	if err != nil {
		return 0, err
	}
	return postID, nil
}

func (s *postService) GetPostByID(postID int) (*models.Post, error) {
	return s.repo.FindPostByID(postID)
}

func (s *postService) GetAllPosts() (map[int]*models.Post, error) {
	return s.repo.GetAllPosts()
}

func (s *postService) UpdateLikeHistory(postID int, likeID int) error {
	return s.repo.UpdateLikeHistory(postID, likeID)
}
