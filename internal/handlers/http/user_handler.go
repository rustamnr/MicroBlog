package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lsmltesting/MicroBlog/internal/dto"
	"github.com/lsmltesting/MicroBlog/internal/service"
)

type UserHTTPHandler struct {
	UserService service.UserService
}

func NewUserHTTPHandler(userService service.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		UserService: userService,
	}
}

func (h *UserHTTPHandler) sendError(w http.ResponseWriter, message string, statusCode int) {
	errUserDTO := dto.ErrorUserDTO{
		Message: message,
		Time:    time.Now(),
	}

	http.Error(w, errUserDTO.ToString(), statusCode)
}

func (h *UserHTTPHandler) UserHandlerRegister(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserDTO

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := userDTO.ValidateForCreate(); err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := h.UserService.CreateUser(userDTO.Username, userDTO.Email, userDTO.Password)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusConflict)
		return
	}

	user, err := h.UserService.GetUserById(userId)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
