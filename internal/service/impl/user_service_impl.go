package service_impl

import (
	"context"
	"github.com/wrlin1218/url_shortener/internal/models"
	"github.com/wrlin1218/url_shortener/internal/repo"
	"github.com/wrlin1218/url_shortener/internal/service"
)

type UserServiceImpl struct {
	UserRepo repo.UserRepo
	LinkRepo repo.LinkRepo
}

func NewUserService(userRepo repo.UserRepo, linkRepo repo.LinkRepo) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepo: userRepo,
		LinkRepo: linkRepo,
	}
}

func (userService *UserServiceImpl) GetUserByUsername(ctx context.Context, username string) (user *models.User, err error) {
	return userService.UserRepo.GetUserByName(ctx, username)
}

func (userService *UserServiceImpl) GetAllLinksByUserName(ctx context.Context, username string) (links []*models.Link, err error) {
	// 1. check if username exist
	user, err := userService.UserRepo.GetUserByName(ctx, username)
	if err != nil {
		return nil, service.UserNotFoundError.Error()
	}
	if user.Name != username {
		return nil, service.UserNotFoundError.Error()
	}

	// 2. query
	return userService.LinkRepo.GetAllLinkByUser(ctx, user.ID), nil
}

func (userService *UserServiceImpl) CreateUser(ctx context.Context, username, password string) error {
	// 1. check if username exist
	exists := userService.CheckUserExists(ctx, username)
	if exists {
		return service.UserAlreadyExistsError.Error()
	}

	// 2. create
	err := userService.UserRepo.CreateUser(ctx, &models.User{
		Name:     username,
		Password: password,
	})
	if err != nil {
		return service.CreateUserFailedError.Error()
	}
	return nil
}

func (userService *UserServiceImpl) CheckUserExists(ctx context.Context, username string) bool {
	_, err := userService.GetUserByUsername(ctx, username)
	if err != nil {
		return false
	}
	return true
}
