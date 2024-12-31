package repo_impl

import (
	"context"
	"github.com/google/uuid"

	"github.com/wrlin1218/url_shortener/internal/dal/kv"
	"github.com/wrlin1218/url_shortener/internal/models"
	"gorm.io/gorm"
)

type LinkRepoImpl struct {
	BaseRepoImpl
}

func NewLinkRepoImpl(DB *gorm.DB, Cache kv.Cache) *LinkRepoImpl {
	return &LinkRepoImpl{
		BaseRepoImpl{DB, Cache},
	}
}

func (impl *LinkRepoImpl) CreateShortLink(ctx context.Context, link *models.Link) error {
	return impl.DB(ctx).Create(link).Error
}

func (impl *LinkRepoImpl) GetShortLinkByCode(ctx context.Context, shortCode string) (*models.Link, error) {
	var link models.Link
	// 1. get from cache
	cacheKey := "shortLink:code:" + shortCode
	err := impl.cache.Get(cacheKey, &link)
	if err == nil {
		return &link, nil
	}

	// 2. get from rdb
	err = impl.DB(ctx).First(&link, "short_code = ?", shortCode).Error
	if err != nil {
		return nil, err
	}
	impl.cache.Set(cacheKey, link, 0)
	return &link, nil
}

func (impl *LinkRepoImpl) DeleteShortLink(ctx context.Context, shortCode string) error {
	return impl.WithTransaction(ctx, func(ctx context.Context) error {
		// 1. delete from db
		err := impl.DB(ctx).Delete(&models.Link{}, "short_code = ?", shortCode).Error
		if err != nil {
			return err
		}

		// 2. delete from cache
		cacheKey := "shortLink:code:" + shortCode
		err = impl.cache.Del(cacheKey)
		if err != nil {
			return err
		}
		return nil
	})
}

func (impl *LinkRepoImpl) GetAllLinkByUser(ctx context.Context, userID uuid.UUID) []*models.Link {
	var links []*models.Link
	impl.DB(ctx).Where("user_id = ?", userID).Find(&links)
	return links
}
