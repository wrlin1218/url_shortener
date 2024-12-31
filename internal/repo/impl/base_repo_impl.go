package repo_impl

import (
	"context"

	"github.com/wrlin1218/url_shortener/internal/dal/kv"
	"github.com/wrlin1218/url_shortener/internal/repo"
	"gorm.io/gorm"
)

type BaseRepoImpl struct {
	db    *gorm.DB
	cache kv.Cache
}

// DB 确定使用的DB：事务db实例 > 默认db实例
func (impl *BaseRepoImpl) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(repo.TxKey).(*gorm.DB)
	if !ok {
		return impl.db.WithContext(ctx)
	}
	return tx.WithContext(ctx)
}

func (impl *BaseRepoImpl) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	// 1. 开启事务
	tx, ok := ctx.Value(repo.TxKey).(*gorm.DB)
	if !ok {
		tx = impl.db.WithContext(ctx).Begin()
		ctx = context.WithValue(ctx, repo.TxKey, tx)
	}

	// 2. 事务后处理
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else if !ok { // 当前函数开启的事务，由当前函数结束； 否则由外部结束
			tx.Commit()
		}
	}()

	// 3. 执行传入的函数
	err = fn(ctx)
	return err
}
