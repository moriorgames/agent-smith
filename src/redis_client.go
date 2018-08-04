package smith

import (
	"github.com/go-redis/redis"
)

type RedisConnectable interface {
	Ping() (string, error)
}

func (r *Redis) Ping() (string, error) {
	return r.client.Ping().Result()
}

type Redis struct {
	client *redis.Client
}

func ConnectToRedis() RedisConnectable {
	redisObj := Redis{}
	redisObj.client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       5,
	})

	return &redisObj
}
