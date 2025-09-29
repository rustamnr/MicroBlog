package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/lsmltesting/MicroBlog/internal/dto"
	"github.com/lsmltesting/MicroBlog/internal/queue"
	"github.com/lsmltesting/MicroBlog/internal/service/like"
)

type LikeHTTPHandler struct {
	LikeService like.LikeService
	LikeQueue   queue.LikeQueue
}

func NewLikeHTTPHandler(likeQueue queue.LikeQueue, likeService like.LikeService) *LikeHTTPHandler {
	return &LikeHTTPHandler{
		LikeService: likeService,
		LikeQueue:   likeQueue,
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

	err = l.LikeQueue.AddLike(userID, postID)
	if err != nil {
		l.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	queueResponseDTO := dto.LikeQueuedResponse{
		Message:   "like was added to queue",
		UserID:    userID,
		PostID:    postID,
		Timestamp: time.Now(),
	}

	b, err := json.MarshalIndent(queueResponseDTO, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (l *LikeHTTPHandler) HandlerGetAllLikes(w http.ResponseWriter, r *http.Request) {
	likes, err := l.LikeService.GetAllLikes()
	if err != nil {
		l.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(likes, " ", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
