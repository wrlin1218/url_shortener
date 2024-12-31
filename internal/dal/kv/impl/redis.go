package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bytedance/sonic"
	"github.com/go-redis/redis/v8"
	"github.com/wrlin1218/url_shortener/internal/dal/kv"
	"github.com/wrlin1218/url_shortener/pkg/logger"
)

type Redis struct {
	client *redis.Client
	ctx    context.Context
}

// RedisInitOptions 除了address，password外的初始化参数
type RedisInitOptions struct {
	Address  string
	Password string
	DB       int
}

// NewClent 初始化
func NewCient(options any) kv.Cache {
	// 1. parse
	initOptions, ok := options.(RedisInitOptions)
	if !ok {
		logger.Fatal("Parse redis init option failed")
	}

	// 2. init
	client := redis.NewClient(&redis.Options{
		Addr:     initOptions.Address,
		Password: initOptions.Password,
		DB:       0,
	})

	// 3. return
	kv.SetCache(&Redis{
		client: client,
		ctx:    context.Background(),
	})
	return kv.GetCache()
}

/**
缓存的基本操作实现
*/

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	bytes, err := sonic.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(r.ctx, key, bytes, expiration).Err()
}

func (r *Redis) Get(key string, dest interface{}) error {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func (r Redis) Del(keys ...string) error {
	return r.client.Del(r.ctx, keys...).Err()
}
