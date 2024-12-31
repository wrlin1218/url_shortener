package repo

import (
	"context"
	"github.com/google/uuid"

	"github.com/wrlin1218/url_shortener/internal/models"
)

type LinkRepo interface {
	BaseRepo
	CreateShortLink(ctx context.Context, link *models.Link) error
	GetShortLinkByCode(ctx context.Context, shortCode string) (*models.Link, error)
	DeleteShortLink(ctx context.Context, shortCode string) error
	GetAllLinkByUser(ctx context.Context, userID uuid.UUID) []*models.Link
	GetAllLink(ctx context.Context) []*models.Link
}
