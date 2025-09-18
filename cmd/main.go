package main

import (
	"time"

	handlers "github.com/lsmltesting/MicroBlog/internal/handlers/http"
	"github.com/lsmltesting/MicroBlog/internal/server"
	"github.com/lsmltesting/MicroBlog/internal/service/post"
	"github.com/lsmltesting/MicroBlog/internal/service/user"
)

func main() {
	userRepo := user.NewInMemoryUserRepo()
	userService := user.NewUserService(userRepo)

	postRepo := post.NewInMemoryPostRepo()
	postService := post.NewPostService(postRepo, userService)

	userHttpHandler := handlers.NewUserHTTPHandler(userService)
	postHttpHandler := handlers.NewPostHTTPHandler(postService, userService)

	serverConfig := server.Config{
		Port:           "8080",
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
	}

	server := server.NewHTTPServer(serverConfig, userHttpHandler, postHttpHandler)
	server.Run()
}
