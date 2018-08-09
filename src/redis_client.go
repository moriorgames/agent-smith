package smith

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisConnectable interface {
	Ping() (string, error)
	Set(string, string, time.Duration) (string, error)
	Get(string) (string, error)
	Del(string) (int64, error)
}

func (r *Redis) Ping() (string, error) {
	return r.client.Ping().Result()
}

func (r *Redis) Set(key string, value string, expr time.Duration) (string, error) {
	return r.client.Set(key, value, expr).Result()
}

func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *Redis) Del(key string) (int64, error) {
	return r.client.Del(key).Result()
}

type Redis struct {
	client *redis.Client
}

func ConnectToRedis() RedisConnectable {
	redisObj := Redis{}
	redisObj.client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       5,
	})

	return &redisObj
}
