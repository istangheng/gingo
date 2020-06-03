package common

import (
	"gingo/config"

	"github.com/go-redis/redis"
)

// RedisClient 实例
var RedisClient *redis.Client

// InitRedis 初始化Redis
func InitRedis() {
	conf := config.Conf

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Pass,
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic("failed to connect redis,err:" + err.Error())
	}

	RedisClient = client
}
