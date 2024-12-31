package kv

import "time"

var cache Cache

func GetCache() Cache {
	return cache
}

func SetCache(cacheImpl Cache) {
	cache = cacheImpl
}

type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string, dest interface{}) error
	Del(keys ...string) error
}
