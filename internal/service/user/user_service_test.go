package user

import (
	"testing"

	"github.com/lsmltesting/MicroBlog/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user *models.User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepository) FindUserByID(ID int) (*models.User, error) {
	args := m.Called(ID)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdatePostHistory(userID int, postID int) error {
	args := m.Called(userID, postID)
	return args.Error(0)
}

func TestUserService_CreateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	testID := 42
	// Set up mock for Save()
	mockRepo.On("Save", mock.MatchedBy(func(user *models.User) bool {
		return user.Username == "testingUser" && user.Email == "test@gmail.com"
	})).Return(testID, nil)

	userId, err := service.CreateUser("testingUser", "test@gmail.com", "sdfmdsfmsdkfm123")

	assert.NoError(t, err)
	assert.Equal(t, testID, userId)

	// Check that method Save is called only 1 time
	mockRepo.AssertExpectations(t)
}

// Check that service work correctly with errors
func TestUserService_CreateUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	//Set up mock for Save() that Save() return error
	mockRepo.On("Save", mock.Anything).Return(0, assert.AnError)

	userId, err := service.CreateUser("testgingUser", "test@gmail.com", "password123")

	assert.Error(t, err)
	assert.Equal(t, 0, userId)

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userID := 33

	expectedUser := &models.User{
		Username: "test_dev_user",
		Email:    "test_dev_user@gmail.com",
		Password: "password_qwerty",
		ID:       userID,
	}

	//Set up mock for FindUserByID()
	mockRepo.On("FindUserByID", userID).Return(expectedUser, nil)

	user, err := service.GetUserByID(userID)

	assert.NoError(t, err)
	assert.Equal(t, user, expectedUser)

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userID := 1231

	mockRepo.On("FindUserByID", mock.Anything).Return((*models.User)(nil), assert.AnError)

	user, err := service.GetUserByID(userID)

	assert.Error(t, err)
	assert.Nil(t, user)

	mockRepo.AssertExpectations(t)
}
