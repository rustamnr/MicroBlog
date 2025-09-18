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
	httpHandler []handlers.HTTPHandler
	config      Config
	httpServer  *http.Server
}

func NewHTTPServer(config Config, httpHandler ...handlers.HTTPHandler) *HTTPServer {
	return &HTTPServer{
		httpHandler: httpHandler,
		config:      config,
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

	for _, register := range s.httpHandler {
		register.RegisterRouters(router)
	}

	s.httpServer.ListenAndServe()

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
