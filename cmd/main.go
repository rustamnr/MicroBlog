package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	handlers "github.com/lsmltesting/MicroBlog/internal/handlers/http"
	"github.com/lsmltesting/MicroBlog/internal/logger"
	"github.com/lsmltesting/MicroBlog/internal/queue"
	likeRepo "github.com/lsmltesting/MicroBlog/internal/repo/like"
	postRepo "github.com/lsmltesting/MicroBlog/internal/repo/post"
	userRepo "github.com/lsmltesting/MicroBlog/internal/repo/user"
	"github.com/lsmltesting/MicroBlog/internal/server"
	likeService "github.com/lsmltesting/MicroBlog/internal/service/like"
	postService "github.com/lsmltesting/MicroBlog/internal/service/post"
	userService "github.com/lsmltesting/MicroBlog/internal/service/user"
)

func main() {
	userRepo := userRepo.NewInMemoryUserRepo()
	baseUserService := userService.NewUserService(userRepo)

	postRepo := postRepo.NewInMemoryPostRepo()
	basePostService := postService.NewPostService(postRepo, baseUserService)

	likeRepo := likeRepo.NewInMemoryLikeRepo()
	baseLikeService := likeService.NewLikeService(likeRepo, baseUserService, basePostService)

	lg := logger.NewLogger(
		logger.LoggerConfig{
			BufferSize: 100,
			Workers:    6,
		},
	)

	userServiceDecorator := userService.NewUserServiceDecorator(baseUserService, lg)
	postServiceDecorator := postService.NewPostServiceDecorator(basePostService, lg)
	likeServiceDecorator := likeService.NewLikeServiceDecorator(baseLikeService, lg)

	likeQueue := queue.NewLikeQueue(
		queue.LikeQueueConfig{
			BufferSize: 100,
			Workers:    6,
		},
		// baseLikeService,
		likeServiceDecorator,
	)

	userHttpHandler := handlers.NewUserHTTPHandler(userServiceDecorator)
	postHttpHandler := handlers.NewPostHTTPHandler(postServiceDecorator, userServiceDecorator)
	likeHttpHandler := handlers.NewLikeHTTPHandler(likeQueue, likeServiceDecorator)

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

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// starting server in goroutine
	serverErr := make(chan error, 1)
	go func() {
		lg.AddLog(
			logger.LevelInfo,
			logger.SourceMain,
			make(map[string]string),
			"Starting server",
		)
		serverErr <- server.Run()
	}()

	select {
	case <-ctx.Done():
		lg.AddLog(
			logger.LevelInfo,
			logger.SourceMain,
			make(map[string]string),
			"Recieved shutdown signal",
		)
	case err := <-serverErr:
		lg.AddLog(
			logger.LevelError,
			logger.SourceMain,
			make(map[string]string),
			fmt.Sprintf("Server error: %v", err),
		)
	}

	lg.AddLog(
		logger.LevelInfo,
		logger.SourceMain,
		make(map[string]string),
		"Shutting down",
	)

	likeQueue.Close()
	lg.Close()

	lg.AddLog(
		logger.LevelInfo,
		logger.SourceMain,
		make(map[string]string),
		"Shutdown complete",
	)
}
