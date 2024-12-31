package repo_impl_test

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	. "github.com/wrlin1218/url_shortener/internal/repo/impl"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	redis "github.com/wrlin1218/url_shortener/internal/dal/kv/impl"
	"github.com/wrlin1218/url_shortener/internal/dal/rdb"
	"github.com/wrlin1218/url_shortener/internal/models"
)

var userRepoImpl *UserRepoImpl

func TestMain(m *testing.M) {
	db := rdb.Init(rdb.RDBOption{
		Dialact: "sqlite",
		DSN:     "test.db",
	})
	cache := redis.NewCient(redis.RedisInitOptions{
		Address:  "localhost:6379",
		Password: "redis-url-shortener-qwenjklqjw6987#sa",
		DB:       0,
	})
	userRepoImpl = NewUserRepoImpl(db, cache)
	m.Run()
}

func TestUserRepoImpl_CreateUser(t *testing.T) {
	repo := userRepoImpl
	ctx := context.Background()

	// 创建测试用户
	user := models.NewUser("Test", "Test123")
	err := repo.CreateUser(ctx, user)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
}

func TestUserRepoImpl_GetUserByID(t *testing.T) {
	repo := userRepoImpl
	ctx := context.Background()

	// 测试从数据库获取
	found, err := repo.GetUserByID(ctx, "5cf2b0d8-a82b-4779-821f-2ea5bfe942f6")
	fmt.Println(sonic.Marshal(found))
	assert.NoError(t, err)

	// 测试从缓存获取
	cached, err := repo.GetUserByID(ctx, "5cf2b0d8-a82b-4779-821f-2ea5bfe942f6")
	fmt.Println(sonic.Marshal(cached))
	assert.NoError(t, err)

	// 测试不存在的用户
	_, err = repo.GetUserByID(ctx, uuid.New().String())
	assert.Error(t, err)
}

func TestUserRepoImpl_GetUserByName(t *testing.T) {
	repo := userRepoImpl
	ctx := context.Background()

	// 测试从数据库获取
	found, err := repo.GetUserByName(ctx, "Test")
	fmt.Println(sonic.Marshal(found))
	assert.NoError(t, err)

	// 测试从缓存获取
	cached, err := repo.GetUserByName(ctx, "Test")
	fmt.Println(sonic.Marshal(cached))
	assert.NoError(t, err)

	// 测试不存在的用户
	_, err = repo.GetUserByName(ctx, "non_existent")
	assert.Error(t, err)
}

func TestUserRepoImpl_DeleteUserByID(t *testing.T) {
	repo := userRepoImpl
	ctx := context.Background()

	// 删除用户
	err := repo.DeleteUserByID(ctx, "5cf2b0d8-a82b-4779-821f-2ea5bfe942f6")
	assert.NoError(t, err)

	// 验证用户已被删除
	_, err = repo.GetUserByID(ctx, "5cf2b0d8-a82b-4779-821f-2ea5bfe942f6")
	assert.Error(t, err)
}

func TestUserRepoImpl_DeleteUserByName(t *testing.T) {
	repo := userRepoImpl
	ctx := context.Background()

	// 删除用户
	err := repo.DeleteUserByName(ctx, "Test2")
	assert.NoError(t, err)

	// 验证用户已被删除
	_, err = repo.GetUserByName(ctx, "Test2")
	assert.Error(t, err)
}
