package main

import (
	"time"

	handlers "github.com/lsmltesting/MicroBlog/internal/handlers/http"
	"github.com/lsmltesting/MicroBlog/internal/server"
	"github.com/lsmltesting/MicroBlog/internal/service"
)

func main() {
	userRepo := service.NewInMemoryUserRepo()
	userService := service.NewUserService(userRepo)

	postRepo := service.NewInMemoryPostRepo()
	postService := service.NewPostService(postRepo, userService)

	userHttpHandler := handlers.NewUserHTTPHandler(userService)
	postHttpHandler := handlers.NewPostHTTPHandler(postService)

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
