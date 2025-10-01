package post

import (
	"fmt"
	"testing"

	"github.com/lsmltesting/MicroBlog/internal/models"
	"github.com/lsmltesting/MicroBlog/internal/repo/post"
)

type mockUserService struct {
	userID int
}

func (m *mockUserService) CreateUser(userName string, email string, password string) (int, error) {
	m.userID++
	return m.userID, nil
}

func (m *mockUserService) GetUserByID(ID int) (*models.User, error) {
	return &models.User{
		Username: "testUserName",
		Email:    "test@test.ru",
		Password: "testPassword",
		ID:       ID,
	}, nil
}

func (m *mockUserService) UpdatePostHistory(userID int, postID int) error {
	return nil
}

func BenchmarkCreatePost(b *testing.B) {
	postRepo := post.NewInMemoryPostRepo()
	mockUserService := &mockUserService{}

	postService := NewPostService(postRepo, mockUserService)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postService.CreatePost(i, fmt.Sprintf("testTextForPost-%d", i))
	}
}

func BenchmarkGetPostByID(b *testing.B) {
	postRepo := post.NewInMemoryPostRepo()
	mockUserService := &mockUserService{}

	postService := NewPostService(postRepo, mockUserService)

	postsID := make([]int, b.N)
	// creating posts
	for i := 0; i < b.N; i++ {
		postID, _ := postService.CreatePost(i, fmt.Sprintf("testTextForPost-%d", i))
		postsID[i] = postID
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postService.GetPostByID(postsID[i])
	}
}

func BenchmarkGetAllPosts(b *testing.B) {
	postRepo := post.NewInMemoryPostRepo()
	mockUserService := &mockUserService{}

	postService := NewPostService(postRepo, mockUserService)

	postsID := make([]int, b.N)
	// creating posts
	for i := 0; i < b.N; i++ {
		postID, _ := postService.CreatePost(i, fmt.Sprintf("testTextForPost-%d", i))
		postsID[i] = postID
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postService.GetAllPosts()
	}
}

func BenchmarkUpdateLikeHistory(b *testing.B) {
	postRepo := post.NewInMemoryPostRepo()
	mockUserService := &mockUserService{}

	postService := NewPostService(postRepo, mockUserService)

	postsID := make([]int, b.N)
	// creating posts
	for i := 0; i < b.N; i++ {
		postID, _ := postService.CreatePost(i, fmt.Sprintf("testTextForPost-%d", i))
		postsID[i] = postID
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postService.UpdateLikeHistory(postsID[i], i)
	}
}
