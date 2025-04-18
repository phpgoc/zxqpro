package interfaces

import (
	"time"

	"github.com/phpgoc/zxqpro/utils"

	"github.com/phpgoc/zxqpro/impl"
)

type CacheImpl byte

type CacheInterface interface {
	Set(key string, value interface{}, expiration time.Duration)
	IsSet(key string) bool
	Get(key string, result interface{}) (res bool)
	Increment(key string, n int64) error
	Decrement(key string, n int64) error
	Delete(key string)
	GetAndRefresh(key string, expiration time.Duration) (interface{}, bool)
}

var Cache CacheInterface = impl.NewGoCache(time.Minute, time.Minute)

func InitCache() {
	if utils.RedisAddr != "" {
		Cache = impl.NewRedisCache(utils.RedisAddr, "", 0)
	} else {
		Cache = impl.NewGoCache(time.Minute, time.Minute)
	}
}
