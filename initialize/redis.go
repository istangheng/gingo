package initialize

import (
	"gingo/internal/config"
	"github.com/garyburd/redigo/redis"
	"time"
)

// Pool redis conn pool
var Pool *redis.Pool

// InitRedisPool 初始化redis conn pool
func InitRedisPool() {
	conf := config.Conf
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", conf.Redis.Addr) },
	}
	Pool = pool
}

// 获取连接
// conn := cache.Pool.Get()
// defer conn.Close()
