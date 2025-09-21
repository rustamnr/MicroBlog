package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/lsmltesting/MicroBlog/internal/dto"
	"github.com/lsmltesting/MicroBlog/internal/service/like"
)

type LikeHTTPHandler struct {
	LikeService like.LikeService
}

func NewLikeHTTPHandler(likeService like.LikeService) *LikeHTTPHandler {
	return &LikeHTTPHandler{
		LikeService: likeService,
	}
}

func (l *LikeHTTPHandler) sendError(w http.ResponseWriter, message string, statusCode int) {
	errUserDTO := dto.ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}

	http.Error(w, errUserDTO.ToString(), statusCode)
}

func (l *LikeHTTPHandler) HandlerCreateLike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postID, err := strconv.Atoi(vars["post_id"])
	if err != nil {
		l.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		l.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	likeID, err := l.LikeService.CreateLike(userID, postID)
	if err != nil {
		l.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	like, err := l.LikeService.GetLikeById(likeID)
	if err != nil {
		l.sendError(w, err.Error(), http.StatusBadRequest)
	}

	b, err := json.MarshalIndent(like, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
