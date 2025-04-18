package interfaces

import (
	"time"

	"github.com/phpgoc/zxqpro/my_runtime"

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
	GetAndRefresh(key string, result interface{}, expiration time.Duration) bool
}

var Cache CacheInterface = impl.NewGoCache(time.Minute, time.Minute)

func InitCache() {
	if my_runtime.RedisAddr != "" {
		Cache = impl.NewRedisCache(my_runtime.RedisAddr, "", 0)
	} else {
		Cache = impl.NewGoCache(time.Minute, time.Minute)
	}
}

func GetOrSet[T any](c CacheInterface, key string, getValue func() T, expiration time.Duration) T {
	var result T
	if c.IsSet(key) {
		if c.GetAndRefresh(key, &result, expiration) {
			return result
		}
	}
	val := getValue()
	c.Set(key, val, expiration)
	return val
}
