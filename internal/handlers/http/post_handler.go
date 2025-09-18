package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/lsmltesting/MicroBlog/internal/dto"
	"github.com/lsmltesting/MicroBlog/internal/service/post"
	"github.com/lsmltesting/MicroBlog/internal/service/user"
)

type PostHTTPHandler struct {
	PostService post.PostService
	UserService user.UserService
}

func NewPostHTTPHandler(postService post.PostService, userService user.UserService) *PostHTTPHandler {
	return &PostHTTPHandler{
		PostService: postService,
		UserService: userService,
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

	postID, err := p.PostService.CreatePost(postDTO.UserID, postDTO.Text)
	if err != nil {
		p.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := p.PostService.GetPostByID(postID)
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

func (p *PostHTTPHandler) HandlerGetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := p.PostService.GetAllPosts()
	if err != nil {
		p.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(posts, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (p *PostHTTPHandler) HandlerAddLikeToPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postID, err := strconv.Atoi(vars["post_id"])
	if err != nil {
		p.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		p.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if post with postID exists
	post, err := p.PostService.GetPostByID(postID)
	if err != nil {
		p.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if user with userID exists
	user, err := p.UserService.GetUserByID(userID)
	if err != nil {
		p.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = p.PostService.AddLikeToPost(user, postID)
	if err != nil {
		p.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(post, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// RegisterRouters registers HTTP routes for handler users
func (p *PostHTTPHandler) RegisterRouters(router *mux.Router) {
	router.Path("/posts").Methods("POST").HandlerFunc(p.HandlerCreatePost)
	router.Path("/posts").Methods("GET").HandlerFunc(p.HandlerGetAllPosts)
	router.Path("/posts/{post_id}/like").Methods("POST").Queries("user_id", "{user_id}").HandlerFunc(p.HandlerAddLikeToPost)
}
