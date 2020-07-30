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
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.Redis.Addr)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", conf.Redis.Pass); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", conf.Redis.Db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
	Pool = pool
}

// 获取连接
// conn := Pool.Get()
// defer conn.Close()

// SetPing 测试
func SetPing() (result string, err error) {
	conn := Pool.Get()
	defer conn.Close()

	n, err := conn.Do("SET", "ping", "pong")
	if err != nil {
		return "", err
	}
	return n.(string), nil
}
