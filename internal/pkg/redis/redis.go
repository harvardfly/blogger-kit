package redis

import (
	"blogger-kit/internal/pkg/config"
	"fmt"
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool
var redisLock *redsync.Redsync

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

func GetRedisConn() (redis.Conn, error) {
	conn := pool.Get()
	return conn, conn.Err()
}
