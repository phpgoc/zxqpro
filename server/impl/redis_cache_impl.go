package impl

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/phpgoc/zxqpro/utils"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(addr string, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	ctx := context.Background()
	err := client.Ping(ctx).Err() // Ping to check if the connection is successful
	if err != nil {
		panic(err)
	}
	return &RedisCache{client: client, ctx: ctx}
}

func (r *RedisCache) Set(key string, value interface{}, expiration time.Duration) {
	data, err := json.Marshal(value)
	if err != nil {
		utils.LogError(err.Error())
		return
	}
	err = r.client.Set(r.ctx, key, data, expiration).Err()
	if err != nil {
		utils.LogError(err.Error())
	}
}

func (r *RedisCache) IsSet(key string) bool {
	return r.client.Exists(r.ctx, key).Val() > 0
}

func (r *RedisCache) Get(key string, result interface{}) (res bool) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		result = nil
		res = false
		return
	}
	err = json.Unmarshal([]byte(val), result)
	if err != nil {
		utils.LogError(err.Error())
		result = nil
		res = false
		return
	}
	res = true
	return
}

func (r *RedisCache) Increment(key string, n int64) error {
	_, err := r.client.IncrBy(r.ctx, key, n).Result()
	return err
}

func (r *RedisCache) Decrement(key string, n int64) error {
	_, err := r.client.DecrBy(r.ctx, key, n).Result()
	return err
}

func (r *RedisCache) Delete(key string) {
	r.client.Del(r.ctx, key)
}

func (r *RedisCache) GetAndRefresh(key string, result interface{}, expiration time.Duration) bool {
	if r.IsSet(key) {
		r.Get(key, result)
		r.client.Expire(r.ctx, key, expiration)
		return true
	} else {
		return false
	}
}
