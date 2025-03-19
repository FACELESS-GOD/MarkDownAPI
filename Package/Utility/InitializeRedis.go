package Utility

import (
	"MarkDownAPI/Helper"
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var RedisInstance *redis.Client

func InitialiseRedisConn() {
	redisOption := &redis.Options{
		Addr:     Helper.RedisAddress,
		Password: Helper.RedisPassword,
		DB:       0,
	}
	RedisInstance = redis.NewClient(redisOption)

}
