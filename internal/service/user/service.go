package user

import "github.com/lsmltesting/MicroBlog/internal/models"

type UserService interface {
	CreateUser(username string, email string, password string) (int, error)
	GetUserByID(ID int) (*models.User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(username string, email string, password string) (int, error) {
	user, err := models.NewUser(username, email, password)
	if err != nil {
		return 0, err
	}
	return s.repo.Save(user)
}

func (s *userService) GetUserByID(ID int) (*models.User, error) {
	return s.repo.FindUserByID(ID)
}
