package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	handlers "github.com/lsmltesting/MicroBlog/internal/handlers/http"
)

type Config struct {
	Port           string
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
}

type HTTPServer struct {
	userHttpHandler *handlers.UserHTTPHandler
	postHttpHandler *handlers.PostHTTPHandler
	likeHttpHandler *handlers.LikeHTTPHandler
	config          Config
	httpServer      *http.Server
}

func NewHTTPServer(
	config Config,
	userHttpHandler *handlers.UserHTTPHandler,
	postHttpHandler *handlers.PostHTTPHandler,
	likeHttpHandler *handlers.LikeHTTPHandler,
) *HTTPServer {
	return &HTTPServer{
		userHttpHandler: userHttpHandler,
		postHttpHandler: postHttpHandler,
		likeHttpHandler: likeHttpHandler,
		config:          config,
	}
}

func (s *HTTPServer) Run() error {
	router := mux.NewRouter()

	s.httpServer = &http.Server{
		Addr:           ":" + s.config.Port,
		Handler:        router,
		MaxHeaderBytes: s.config.MaxHeaderBytes,
		ReadTimeout:    s.config.ReadTimeout,
		WriteTimeout:   s.config.WriteTimeout,
		IdleTimeout:    s.config.IdleTimeout,
	}

	// Register methods from userHttpHandler
	router.Path("/register").Methods("POST").HandlerFunc(s.userHttpHandler.UserHandlerRegister)

	// Register methods from postHttpHandler
	router.Path("/posts").Methods("POST").HandlerFunc(s.postHttpHandler.HandlerCreatePost)
	router.Path("/posts").Methods("GET").HandlerFunc(s.postHttpHandler.HandlerGetAllPosts)

	// Register method from likeHttpHandler
	router.Path("/posts/{post_id}/like").Methods("POST").Queries("user_id", "{user_id}").HandlerFunc(s.likeHttpHandler.HandlerCreateLike)

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
