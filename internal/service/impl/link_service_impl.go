package service_impl

import (
	"context"
	"github.com/spaolacci/murmur3"
	"github.com/wrlin1218/url_shortener/internal/models"
	"github.com/wrlin1218/url_shortener/internal/repo"
	"github.com/wrlin1218/url_shortener/internal/service"
	"strconv"
)

type LinkServiceImpl struct {
	LinkRepo repo.LinkRepo
	UserRepo repo.UserRepo
}

func NewLinkService(linkRepo repo.LinkRepo, userRepo repo.UserRepo) *LinkServiceImpl {
	return &LinkServiceImpl{
		LinkRepo: linkRepo,
		UserRepo: userRepo,
	}
}

func (linkService *LinkServiceImpl) CreateShortLink(ctx context.Context, username string, OriginalUrl string) (err error, shortCode string) {
	// 1. generate shortcode
	user, err := linkService.UserRepo.GetUserByName(ctx, username)
	if err != nil {
		return service.UserNotFoundError.Error(), ""
	}

	hash := murmur3.Sum32([]byte(user.ID.String() + OriginalUrl))
	shortCode = strconv.FormatUint(uint64(hash), 16)

	// 2. generate Link model
	err = linkService.LinkRepo.CreateShortLink(ctx, &models.Link{
		UserID:      user.ID,
		OriginalURL: OriginalUrl,
		ShortCode:   shortCode,
	})
	return err, shortCode
}

func (linkService *LinkServiceImpl) GetOriginalUrl(ctx context.Context, shortCode string) (err error, originalUrl string) {
	code, err := linkService.LinkRepo.GetShortLinkByCode(ctx, shortCode)
	if err != nil {
		return err, ""
	}
	return nil, code.OriginalURL
}

func (linkService *LinkServiceImpl) DeleteShortLink(ctx context.Context, username string, shortCode string) (err error) {
	user, err := linkService.UserRepo.GetUserByName(ctx, username)
	if err != nil {
		return service.UserNotFoundError.Error()
	}

	// 1. get link
	code, err := linkService.LinkRepo.GetShortLinkByCode(ctx, shortCode)
	if err != nil {
		return service.LinkNotFoundError.Error()
	}

	// 2. check if link belongs to user
	if code.UserID != user.ID {
		return service.NoPermissionToOperateError.Error()
	}

	// 3. delete
	return linkService.LinkRepo.DeleteShortLink(ctx, shortCode)
}
