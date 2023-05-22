package cache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/ilhamnyto/echo-fw/config"
)


func ConnectRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.GetString("REDIS_HOST"), config.GetString("REDIS_PORT")),
		Password: config.GetString("REDIS_PASSWORD"),
		DB: config.GetInt("REDIS_DB"),
	})

	return redisClient
}