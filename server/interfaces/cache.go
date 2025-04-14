package interfaces

import (
	"time"

	"github.com/phpgoc/zxqpro/impl"
)

type CacheImpl byte

const (
	CacheGoImpl CacheImpl = iota
	CacheRedisImpl
)

type CacheInterface interface {
	Set(key string, value interface{}, expiration time.Duration)
	IsSet(key string) bool
	Get(key string, result interface{}) (res bool)
	Increment(key string, n int64) error
	Decrement(key string, n int64) error
	Delete(key string)
	GetAndRefresh(key string, expiration time.Duration) (interface{}, bool)
}

var Cache = cacheFactory(CacheRedisImpl)

func cacheFactory(i CacheImpl) CacheInterface {
	switch i {
	case CacheGoImpl:
		return impl.NewGoCache(time.Minute, time.Minute)
	case CacheRedisImpl:
		return impl.NewRedisCache("localhost:6379", "", 0)
	default:
		return impl.NewGoCache(time.Minute, time.Minute)
	}
}
