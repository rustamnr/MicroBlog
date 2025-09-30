package user

import (
	"strconv"

	"github.com/lsmltesting/MicroBlog/internal/logger"
	"github.com/lsmltesting/MicroBlog/internal/models"
)

type userServiceDecorator struct {
	userService UserService
	lg          logger.Logger
}

func NewUserServiceDecorator(userService UserService, lg logger.Logger) UserService {
	return &userServiceDecorator{
		userService: userService,
		lg:          lg,
	}
}

func (u *userServiceDecorator) CreateUser(username string, email string, password string) (int, error) {
	u.lg.AddLog(
		logger.LevelInfo,
		logger.SourceService,
		map[string]string{
			"username": username,
			"email":    email,
			"password": password,
		},
		"CreateUser from like service is called",
	)

	userID, err := u.userService.CreateUser(username, email, password)

	if err != nil {
		u.lg.AddLog(
			logger.LevelError,
			logger.SourceService,
			map[string]string{
				"username": username,
				"email":    email,
				"password": password,
			},
			err.Error(),
		)
	} else {
		u.lg.AddLog(
			logger.LevelInfo,
			logger.SourceService,
			map[string]string{
				"username": username,
				"email":    email,
				"password": password,
			},
			"CreateUser called successfully",
		)
	}

	return userID, err
}
func (u *userServiceDecorator) GetUserByID(ID int) (*models.User, error) {
	u.lg.AddLog(
		logger.LevelInfo,
		logger.SourceService,
		map[string]string{
			"ID": strconv.Itoa(ID),
		},
		"GetUserByID from like service is called",
	)

	userModel, err := u.userService.GetUserByID(ID)

	if err != nil {
		u.lg.AddLog(
			logger.LevelError,
			logger.SourceService,
			map[string]string{
				"ID": strconv.Itoa(ID),
			},
			err.Error(),
		)
	} else {
		u.lg.AddLog(
			logger.LevelInfo,
			logger.SourceService,
			map[string]string{
				"ID": strconv.Itoa(ID),
			},
			"GetUserByID called successfully",
		)
	}

	return userModel, err
}
func (u *userServiceDecorator) UpdatePostHistory(userID int, postID int) error {
	u.lg.AddLog(
		logger.LevelInfo,
		logger.SourceService,
		map[string]string{
			"userID": strconv.Itoa(userID),
			"postID": strconv.Itoa(postID),
		},
		"UpdatePostHistory from like service is called",
	)

	err := u.userService.UpdatePostHistory(userID, postID)

	if err != nil {
		u.lg.AddLog(
			logger.LevelError,
			logger.SourceService,
			map[string]string{
				"userID": strconv.Itoa(userID),
				"postID": strconv.Itoa(postID),
			},
			err.Error(),
		)
	} else {
		u.lg.AddLog(
			logger.LevelInfo,
			logger.SourceService,
			map[string]string{
				"userID": strconv.Itoa(userID),
				"postID": strconv.Itoa(postID),
			},
			"UpdatePostHistory called successfully",
		)
	}

	return err
}
