package post

import (
	"testing"

	"github.com/lsmltesting/MicroBlog/internal/models"
	"github.com/lsmltesting/MicroBlog/internal/service/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPostRepository struct {
	mock.Mock
}

type MockUserRepository struct {
	mock.Mock
}

// Posts methods
func (mockPost *MockPostRepository) Save(post *models.Post) (int, error) {
	args := mockPost.Called(post)
	return args.Int(0), args.Error(1)
}

func (mockPost *MockPostRepository) FindPostByID(postID int) (*models.Post, error) {
	args := mockPost.Called(postID)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (mockPost *MockPostRepository) GetAllPosts() (map[int]*models.Post, error) {
	args := mockPost.Called()
	return args.Get(0).(map[int]*models.Post), args.Error(1)
}

func (mockPost *MockPostRepository) AddLikeToPost(postID int, likeID int) error {
	args := mockPost.Called(postID, likeID)
	return args.Error(0)
}

func (mockPost *MockPostRepository) UpdateLikeHistory(postID int, likeID int) error {
	args := mockPost.Called(postID, likeID)
	return args.Error(0)
}

// Users methods
func (mockUser *MockUserRepository) Save(user *models.User) (int, error) {
	args := mockUser.Called(user)
	return args.Int(0), args.Error(1)
}

func (mockUser *MockUserRepository) FindUserByID(ID int) (*models.User, error) {
	args := mockUser.Called(ID)
	return args.Get(0).(*models.User), args.Error(1)
}

func (mockUser *MockUserRepository) UpdatePostHistory(userID int, postID int) error {
	args := mockUser.Called(userID, postID)
	return args.Error(0)
}

func TestPostService_Save_Success(t *testing.T) {
	mockRepoPost := new(MockPostRepository)
	mockRepoUser := new(MockUserRepository)

	userService := user.NewUserService(mockRepoUser)
	postService := NewPostService(mockRepoPost, userService)

	postID := 2
	textPost := "smth new test text for post is already written"

	userID := 1

	testUser := &models.User{
		Username: "test_user_name_for_post",
		Email:    "test_mail@gmail.com",
		Password: "test_password_qwerty_for_post",
		ID:       userID,
	}

	mockRepoUser.On("FindUserByID", userID).Return(testUser, nil)
	mockRepoPost.On("Save", mock.MatchedBy(func(post *models.Post) bool {
		return (post.Text == textPost &&
			post.UserID == userID)
	})).Return(postID, nil)

	mockRepoUser.On("UpdatePostHistory", userID, postID).Return(nil)

	createdPostId, err := postService.CreatePost(userID, textPost)

	assert.NoError(t, err)
	assert.Equal(t, postID, createdPostId)

	mockRepoPost.AssertExpectations(t)
	mockRepoUser.AssertExpectations(t)
}

func TestPostService_Save_Error(t *testing.T) {
	mockPostRepository := new(MockPostRepository)
	mockUserRepository := new(MockUserRepository)

	userService := user.NewUserService(mockUserRepository)
	postService := NewPostService(mockPostRepository, userService)

	userID := 10
	postText := "test post text"

	testUser := &models.User{
		Username: "test_user_name_for_post",
		Email:    "test_mail@gmail.com",
		Password: "test_password_qwerty_for_post",
		ID:       userID,
	}

	mockUserRepository.On("FindUserByID", userID).Return(testUser, nil)
	mockPostRepository.On("Save", mock.Anything).Return(0, assert.AnError)
	mockUserRepository.On("UpdatePostHistory", userID, mock.Anything).Return(nil).Maybe()

	postID, err := postService.CreatePost(userID, postText)

	assert.Error(t, err)
	assert.Equal(t, 0, postID)

	mockPostRepository.AssertExpectations(t)
	mockUserRepository.AssertExpectations(t)
}

func TestPostService_FindPostByID_Success(t *testing.T) {
	mockPostRepository := new(MockPostRepository)
	mockUserRepository := new(MockUserRepository)

	userService := user.NewUserService(mockUserRepository)
	postService := NewPostService(mockPostRepository, userService)

	testPostID := 13
	testUserID := 7
	testPostText := "random text for post test"

	expectedPost := &models.Post{
		Text:   testPostText,
		UserID: testUserID,
		// User: &models.User{
		// 	Username: "test_user",
		// 	Email:    "test_user_for_post@gmail.com",
		// 	Password: "qwert123password",
		// 	ID:       10,
		// },
	}

	mockPostRepository.On("FindPostByID", testPostID).Return(expectedPost, nil)

	post, err := postService.GetPostByID(testPostID)

	assert.NoError(t, err)
	assert.Equal(t, post, expectedPost)

	mockPostRepository.AssertExpectations(t)
	mockUserRepository.AssertExpectations(t)
}

func TestPostService_FindPostByID_NotFound(t *testing.T) {
	mockPostRepository := new(MockPostRepository)
	mockUserRepository := new(MockUserRepository)

	userService := user.NewUserService(mockUserRepository)
	postService := NewPostService(mockPostRepository, userService)

	testPostID := 13

	mockPostRepository.On("FindPostByID", mock.Anything).Return((*models.Post)(nil), assert.AnError)

	post, err := postService.GetPostByID(testPostID)

	assert.Error(t, err)
	assert.Nil(t, post)

	mockPostRepository.AssertExpectations(t)
	mockUserRepository.AssertExpectations(t)
}

func TestPostService_GetAllPosts_Success(t *testing.T) {
	mockPostRepository := new(MockPostRepository)
	mockUserRepository := new(MockUserRepository)

	userService := user.NewUserService(mockUserRepository)
	postService := NewPostService(mockPostRepository, userService)

	expectedPosts := map[int]*models.Post{
		1: {
			ID:     10,
			Text:   "first test text",
			UserID: 9,
			// User: &models.User{
			// 	Username: "first temp user",
			// 	Email:    "firstemailtempusre@mail.ru",
			// 	Password: "testqwerty123",
			// },
		},
		2: {
			ID:     33,
			Text:   "second test text",
			UserID: 123,
			// User: &models.User{
			// 	Username: "second temp user",
			// 	Email:    "secondemailtempusre@mail.ru",
			// 	Password: "second_testqwerty123",
			// },
		},
	}

	mockPostRepository.On("GetAllPosts").Return(expectedPosts, nil)

	posts, err := postService.GetAllPosts()

	assert.NoError(t, err)
	assert.Equal(t, expectedPosts, posts)

	mockPostRepository.AssertExpectations(t)
	mockUserRepository.AssertExpectations(t)
}

func TestPostService_GetAllPosts_Error(t *testing.T) {
	mockPostRepository := new(MockPostRepository)
	mockUserRepository := new(MockUserRepository)

	userService := user.NewUserService(mockUserRepository)
	postService := NewPostService(mockPostRepository, userService)

	mockPostRepository.On("GetAllPosts").Return((map[int]*models.Post)(nil), assert.AnError)

	posts, err := postService.GetAllPosts()

	assert.Error(t, err)
	assert.Nil(t, posts)

	mockPostRepository.AssertExpectations(t)
	mockUserRepository.AssertExpectations(t)
}

// func TestPostService_AddLikeToPost_Success(t *testing.T) {
// 	mockPostRepository := new(MockPostRepository)
// 	mockUserRepository := new(MockUserRepository)

// 	userService := user.NewUserService(mockUserRepository)
// 	postService := NewPostService(mockPostRepository, userService)

// 	expectedUser := &models.User{
// 		Username: "temp user",
// 		Email:    "temp_test_user@gmail.com",
// 		Password: "qw123ery_password",
// 	}
// 	expectedPostID := 10

// 	mockPostRepository.On("AddLikeToPost",
// 		mock.MatchedBy(func(user *models.User) bool {
// 			return user.Username == expectedUser.Username &&
// 				user.Email == expectedUser.Email &&
// 				user.Password == expectedUser.Password
// 		}),
// 		expectedPostID,
// 		mock.MatchedBy(func(like *models.Like) bool {
// 			return like != nil &&
// 				like.UserFrom.ID == expectedUser.ID
// 		}),
// 	).Return(nil)

// 	err := postService.AddLikeToPost(expectedUser, expectedPostID)

// 	assert.NoError(t, err)

// 	mockPostRepository.AssertExpectations(t)
// 	mockUserRepository.AssertExpectations(t)
// }

// func TestPostService_AddLikeToPost_Error(t *testing.T) {
// 	mockPostRepository := new(MockPostRepository)
// 	mockUserRepository := new(MockUserRepository)

// 	userService := user.NewUserService(mockUserRepository)
// 	postService := NewPostService(mockPostRepository, userService)

// 	expectedUser := &models.User{
// 		Username: "temp user",
// 		Email:    "temp_test_user@gmail.com",
// 		Password: "qw123ery_password",
// 	}
// 	expectedPostID := 10

// 	mockPostRepository.On("AddLikeToPost",
// 		mock.Anything,
// 		mock.Anything,
// 		mock.Anything,
// 	).Return(assert.AnError)

// 	err := postService.AddLikeToPost(expectedUser, expectedPostID)

// 	assert.Error(t, err)

// 	mockPostRepository.AssertExpectations(t)
// 	mockUserRepository.AssertExpectations(t)
// }
