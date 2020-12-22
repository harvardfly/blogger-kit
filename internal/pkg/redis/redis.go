package redis

import (
	"fmt"
	"time"

	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/config"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool
var redisLock *redsync.Redsync

// InitRedis 初始化redis
func InitRedis(cfg *config.RedisConfig) error {
	pool = &redis.Pool{
		MaxIdle:     20,
		IdleTimeout: 240 * time.Second,
		MaxActive:   50,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			)
			if err != nil {
				return nil, err
			}
			if cfg.Password != "" {
				if _, err := c.Do("AUTH", cfg.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	redisLock = redsync.New([]redsync.Pool{pool})
	return nil
}

// GetRedisConn 获取redis连接
func GetRedisConn() (redis.Conn, error) {
	conn := pool.Get()
	return conn, conn.Err()
}
