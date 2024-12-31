package repo_impl

import (
	"context"
	"fmt"

	"github.com/wrlin1218/url_shortener/internal/dal/kv"
	"github.com/wrlin1218/url_shortener/internal/models"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	BaseRepoImpl
}

func NewUserRepoImpl(DB *gorm.DB, Cache kv.Cache) *UserRepoImpl {
	return &UserRepoImpl{
		BaseRepoImpl{DB, Cache},
	}
}

func (impl *UserRepoImpl) CreateUser(ctx context.Context, user *models.User) error {
	return impl.DB(ctx).Create(user).Error
}

func (impl *UserRepoImpl) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	// 1. get from cache
	cacheKey := "user:id:" + id
	err := impl.cache.Get(cacheKey, &user)
	if err == nil {
		return &user, nil
	}

	// 2. get from db
	err = impl.DB(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	impl.cache.Set(cacheKey, user, 0)
	impl.cache.Set("user:name:"+user.Name, user, 0)
	return &user, nil
}

func (impl *UserRepoImpl) GetUserByName(ctx context.Context, name string) (*models.User, error) {
	var user models.User
	// 1. get from cache
	cacheKey := "user:name:" + name
	err := impl.cache.Get(cacheKey, &user)
	if err == nil {
		return &user, nil
	}

	// 2. get from db
	err = impl.DB(ctx).First(&user, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	impl.cache.Set(cacheKey, user, 0)
	impl.cache.Set("user:id:"+user.ID.String(), user, 0)
	return &user, nil
}

func (impl *UserRepoImpl) DeleteUserByID(ctx context.Context, id string) error {
	// 1. 先查询
	user, err := impl.GetUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error user check before delete")
	}

	// 2. 再删除
	err = impl.WithTransaction(ctx, func(ctx context.Context) error {
		// delete form db
		err = impl.DB(ctx).Delete(&models.User{}, "id = ?", id).Error
		if err != nil {
			return err
		}

		// delete from cache
		idCacheKey := "user:id:" + user.ID.String()
		nameCacheKey := "user:name:" + user.Name
		err := impl.cache.Del(idCacheKey, nameCacheKey)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (impl *UserRepoImpl) DeleteUserByName(ctx context.Context, name string) error {
	// 1. 先查询
	user, err := impl.GetUserByName(ctx, name)
	if err != nil {
		return fmt.Errorf("error user check before delete")
	}

	// 2. 再删除
	return impl.DeleteUserByID(ctx, user.ID.String())
}
