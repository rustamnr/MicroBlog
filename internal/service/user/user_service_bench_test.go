package user

import (
	"fmt"
	"testing"

	repo "github.com/lsmltesting/MicroBlog/internal/repo/user"
)

func BenchmarkCreateUser(b *testing.B) {
	userRepo := repo.NewInMemoryUserRepo()
	userService := NewUserService(userRepo)

	users := make(
		[]struct {
			username string
			email    string
			password string
		},
		b.N,
	)
	for i := range users {
		users[i].username = fmt.Sprintf("testUsername-%d", i)
		users[i].email = fmt.Sprintf("testemail-%d@gtest.ru", i)
		users[i].password = fmt.Sprintf("testPassword123qwerty-%d", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		userService.CreateUser(users[i].username, users[i].email, users[i].password)
	}
}
