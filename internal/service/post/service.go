package post

import (
	"github.com/lsmltesting/MicroBlog/internal/models"
	"github.com/lsmltesting/MicroBlog/internal/service/user"
)

type PostService interface {
	CreatePost(user int, text string) (int, error)
	GetPostById(postId int) (*models.Post, error)
	GetAllPosts() (map[int]*models.Post, error)
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

func (s *postService) CreatePost(userId int, text string) (int, error) {
	// Check if user with shared userId is exists
	user, err := s.userService.GetUserById(userId)
	if err != nil {
		return 0, err
	}

	post, err := models.NewPost(user, text)

	if err != nil {
		return 0, err
	}

	return s.repo.Save(post)
}

func (s *postService) GetPostById(postId int) (*models.Post, error) {
	return s.repo.FindPostById(postId)
}

func (s *postService) GetAllPosts() (map[int]*models.Post, error) {
	return s.repo.GetAllPosts()
}
