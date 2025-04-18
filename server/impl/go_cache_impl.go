package impl

import (
	"fmt"
	"reflect"
	"time"

	"github.com/patrickmn/go-cache"
)

// GoCache 封装 go-cache 的结构体
type GoCache struct {
	cache *cache.Cache
}

// NewGoCache 创建一个新的 GoCache 实例
func NewGoCache(defaultExpiration, cleanupInterval time.Duration) *GoCache {
	return &GoCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

// Set 实现 Cache 接口的 Set 方法
func (g *GoCache) Set(key string, value interface{}, expiration time.Duration) {
	g.cache.Set(key, value, expiration)
}

func (g *GoCache) IsSet(key string) bool {
	_, found := g.cache.Get(key)
	return found
}

// Get 实现 Cache 接口的 Get 方法
func (g *GoCache) Get(key string, result interface{}) (res bool) {
	value, found := g.cache.Get(key)
	if found {
		resultValue := reflect.ValueOf(result)
		// 检查 result 是否为指针类型
		if resultValue.Kind() != reflect.Ptr || resultValue.IsNil() {
			fmt.Println("result must be a non-nil pointer")
			return false
		}
		// 获取指针指向的值
		resultElem := resultValue.Elem()
		// 检查缓存值和 result 指向的值的类型是否匹配
		valueType := reflect.TypeOf(value)
		if resultElem.Type() != valueType {
			fmt.Printf("missmatch type，cache type is  %v，expect is %v\n", valueType, resultElem.Type())
			return false
		}
		// 将缓存值赋给 result 指向的内存地址
		resultElem.Set(reflect.ValueOf(value))
		return true
	}
	return false
}

// Increment 实现 Cache 接口的 Increment 方法
func (g *GoCache) Increment(key string, n int64) error {
	return g.cache.Increment(key, n)
}

// Decrement 实现 Cache 接口的 Decrement 方法
func (g *GoCache) Decrement(key string, n int64) error {
	return g.cache.Decrement(key, n)
}

// Delete 实现 Cache 接口的 Delete 方法
func (g *GoCache) Delete(key string) {
	g.cache.Delete(key)
}

func (g *GoCache) GetAndRefresh(key string, expiration time.Duration) (interface{}, bool) {
	value, found := g.cache.Get(key)
	if found {
		// 如果找到值，使用 Set 方法刷新过期时间
		g.cache.Set(key, value, expiration)
	}
	return value, found
}
