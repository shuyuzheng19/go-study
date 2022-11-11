package config

import "github.com/go-redis/redis"

const (
	redisAddr = "127.0.0.1:6379"
)

func GetRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   2,
	})
	return client
}
