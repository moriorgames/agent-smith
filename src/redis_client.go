package smith

import (
	"github.com/go-redis/redis"
)

type RedisConnectable interface {
	Ping() (string, error)
}

type Redis struct {
	client *redis.Client
}

func (r *Redis) Ping() (string, error) {
	return r.client.Ping().Result()
}

func ConnectToRedis(r *Redis) (string, error) {
	r.client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       5,
	})

	return r.client.Ping().Result()
}
