package redis

import (
	"gorm-study/config"
	"time"
)

var redis = config.GetRedis()

func Set(key string, value string, duration time.Duration) bool {
	err := redis.Set(key, value, duration).Err()

	if err != nil {
		return false
	}

	return true

}

func GetString(key string) string {
	result, err := redis.Get(key).Result()

	if err != nil {
		panic(err)
	}

	return result

}
