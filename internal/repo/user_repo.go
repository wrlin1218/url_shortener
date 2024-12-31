package repo

import (
	"context"

	"github.com/wrlin1218/url_shortener/internal/models"
)

type UserRepo interface {
	BaseRepo
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByName(ctx context.Context, name string) (*models.User, error)
	DeleteUserByID(ctx context.Context, id string) error
	DeleteUserByName(ctx context.Context, name string) error
}
