package common

import (
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

// RedisClient 实例
var RedisClient *redis.Client

// InitRedis 初始化Redis
func InitRedis() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       0, // use default DB
	})

	_, err = client.Ping().Result()
	if err != nil {
		panic("failed to connect redis,err:" + err.Error())
	}

	RedisClient = client

}
