package post

import (
	"strconv"

	"github.com/lsmltesting/MicroBlog/internal/logger"
	"github.com/lsmltesting/MicroBlog/internal/models"
)

type postServiceDecorator struct {
	postService PostService
	lg          logger.Logger
}

func NewPostServiceDecorator(postService PostService, lg logger.Logger) PostService {
	return &postServiceDecorator{
		postService: postService,
		lg:          lg,
	}
}

func (p *postServiceDecorator) CreatePost(userID int, text string) (int, error) {
	p.lg.AddLog(
		logger.LevelInfo,
		logger.SourceService,
		map[string]string{
			"userID": strconv.Itoa(userID),
			"text":   text,
		},
		"CreatePost from post service is called",
	)

	postID, err := p.postService.CreatePost(userID, text)
	if err != nil {
		p.lg.AddLog(
			logger.LevelError,
			logger.SourceService,
			map[string]string{
				"userID": strconv.Itoa(userID),
				"text":   text,
				"postID": strconv.Itoa(postID),
			},
			err.Error(),
		)
	} else {
		p.lg.AddLog(
			logger.LevelInfo,
			logger.SourceService,
			map[string]string{
				"userID": strconv.Itoa(userID),
				"text":   text,
				"postID": strconv.Itoa(postID),
			},
			"CreatePost called successfully",
		)
	}

	return postID, err
}

func (p *postServiceDecorator) GetPostByID(postID int) (*models.Post, error) {
	p.lg.AddLog(
		logger.LevelInfo,
		logger.SourceService,
		map[string]string{
			"postID": strconv.Itoa(postID),
		},
		"GetPostByID from post service is called",
	)

	postModel, err := p.postService.GetPostByID(postID)

	if err != nil {
		p.lg.AddLog(
			logger.LevelError,
			logger.SourceService,
			map[string]string{
				"postID": strconv.Itoa(postID),
			},
			err.Error(),
		)
	} else {
		p.lg.AddLog(
			logger.LevelInfo,
			logger.SourceService,
			map[string]string{
				"postID": strconv.Itoa(postID),
			},
			"GetPostByID called successfully",
		)
	}

	return postModel, err
}

func (p *postServiceDecorator) GetAllPosts() (map[int]*models.Post, error) {
	p.lg.AddLog(
		logger.LevelInfo,
		logger.SourceService,
		make(map[string]string),
		"GetAllPosts from post service is called",
	)

	postModel, err := p.postService.GetAllPosts()

	if err != nil {
		p.lg.AddLog(
			logger.LevelError,
			logger.SourceService,
			make(map[string]string),
			err.Error(),
		)
	} else {
		p.lg.AddLog(
			logger.LevelInfo,
			logger.SourceService,
			make(map[string]string),
			"GetAllPosts called successfully",
		)
	}

	return postModel, err
}

func (p *postServiceDecorator) UpdateLikeHistory(postID int, likeID int) error {
	p.lg.AddLog(
		logger.LevelInfo,
		logger.SourceService,
		make(map[string]string),
		"UpdateLikeHistory from post service is called",
	)

	err := p.postService.UpdateLikeHistory(postID, likeID)

	if err != nil {
		p.lg.AddLog(
			logger.LevelError,
			logger.SourceService,
			make(map[string]string),
			err.Error(),
		)
	} else {
		p.lg.AddLog(
			logger.LevelInfo,
			logger.SourceService,
			make(map[string]string),
			"UpdateLikeHistory called successfully",
		)
	}

	return err
}
