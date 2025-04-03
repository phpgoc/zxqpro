package interfaces

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/phpgoc/zxqpro/impl"
)

type CacheInterface interface {
	Set(key string, value interface{}, expiration time.Duration)
	Get(key string) (interface{}, bool)
	Increment(key string, n int64) error
	Decrement(key string, n int64) error
	Delete(key string)
	GetAndRefresh(key string, expiration time.Duration) (interface{}, bool)
}

var Cache CacheInterface = impl.NewGoCache(cache.DefaultExpiration, time.Minute)
