package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lsmltesting/MicroBlog/internal/dto"
	"github.com/lsmltesting/MicroBlog/internal/service"
)

type PostHTTPHandler struct {
	PostService service.PostService
}

func NewPostHTTPHandler(postService service.PostService) *PostHTTPHandler {
	return &PostHTTPHandler{
		PostService: postService,
	}
}

func (h *PostHTTPHandler) sendError(w http.ResponseWriter, message string, statusCode int) {
	errUserDTO := dto.ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}

	http.Error(w, errUserDTO.ToString(), statusCode)
}

func (p *PostHTTPHandler) HandlerCreatePost(w http.ResponseWriter, r *http.Request) {
	var postDTO dto.PostDTO

	// Check to get PostDTO from request body
	if err := json.NewDecoder(r.Body).Decode(&postDTO); err != nil {
		p.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	postId, err := p.PostService.CreatePost(postDTO.UserId, postDTO.Text)
	if err != nil {
		p.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := p.PostService.GetPostById(postId)
	if err != nil {
		p.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(post, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

// RegisterRouters registers HTTP routes for handler users
func (p *PostHTTPHandler) RegisterRouters(router *mux.Router) {
	router.Path("/posts").Methods("POST").HandlerFunc(p.HandlerCreatePost)
}
