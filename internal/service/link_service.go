package service

import (
	"context"
)

type LinkService interface {
	CreateShortLink(ctx context.Context, username string, OriginalUrl string) (err error, shortCode string)
	GetOriginalUrl(ctx context.Context, shortCode string) (err error, originalUrl string)
	DeleteShortLink(ctx context.Context, username string, shortCode string) (err error)
}
