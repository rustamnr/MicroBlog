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
	config          Config
	httpServer      *http.Server
}

func NewHTTPServer(userHttpHandler *handlers.UserHTTPHandler, config Config) *HTTPServer {
	return &HTTPServer{
		userHttpHandler: userHttpHandler,
		config:          config,
	}
}

func (s *HTTPServer) Run() error {
	router := mux.NewRouter()
	router.Path("/register").Methods("POST").HandlerFunc(s.userHttpHandler.UserHandlerRegister)

	s.httpServer = &http.Server{
		Addr:           ":" + s.config.Port,
		Handler:        router,
		MaxHeaderBytes: s.config.MaxHeaderBytes,
		ReadTimeout:    s.config.ReadTimeout,
		WriteTimeout:   s.config.WriteTimeout,
		IdleTimeout:    s.config.IdleTimeout,
	}

	s.httpServer.ListenAndServe()

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
