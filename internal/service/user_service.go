package service

import (
	"context"
	"github.com/wrlin1218/url_shortener/internal/models"
)

type UserService interface {
	GetUserByUsername(ctx context.Context, username string) (user *models.User, err error)
	GetAllLinksByUserName(ctx context.Context, username string) (links []*models.Link, err error)
	CreateUser(ctx context.Context, username, password string) error
	CheckUserExists(ctx context.Context, username string) bool
}
