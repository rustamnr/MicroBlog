package main

import (
	"time"

	handlers "github.com/lsmltesting/MicroBlog/internal/handlers/http"
	"github.com/lsmltesting/MicroBlog/internal/queue"
	"github.com/lsmltesting/MicroBlog/internal/server"
	"github.com/lsmltesting/MicroBlog/internal/service/like"
	"github.com/lsmltesting/MicroBlog/internal/service/post"
	"github.com/lsmltesting/MicroBlog/internal/service/user"
)

func main() {
	userRepo := user.NewInMemoryUserRepo()
	userService := user.NewUserService(userRepo)

	postRepo := post.NewInMemoryPostRepo()
	postService := post.NewPostService(postRepo, userService)

	likeRepo := like.NewInMemoryLikeRepo()
	likeService := like.NewLikeService(likeRepo, userService, postService)

	likeQueue := queue.NewLikeQueue(
		queue.LikeQueueConfig{
			BufferSize: 100,
			Workers:    6,
		},
		likeService,
	)

	userHttpHandler := handlers.NewUserHTTPHandler(userService)
	postHttpHandler := handlers.NewPostHTTPHandler(postService, userService)
	likeHttpHandler := handlers.NewLikeHTTPHandler(likeQueue, likeService)
	defer likeQueue.Close()

	serverConfig := server.Config{
		Port:           "8080",
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
	}

	server := server.NewHTTPServer(
		serverConfig,
		userHttpHandler,
		postHttpHandler,
		likeHttpHandler,
	)
	server.Run()
}
