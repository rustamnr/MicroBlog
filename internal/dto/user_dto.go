package dto

import (
	"encoding/json"
	"time"

	"github.com/lsmltesting/MicroBlog/internal/models"
)

type UserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserDTO) ValidateForCreate() error {
	if _, err := models.NewUser(u.Username, u.Email, u.Email); err != nil {
		return err
	}
	return nil
}

type ErrorUserDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorUserDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
