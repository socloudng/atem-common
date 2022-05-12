package cache

import (
	"fmt"
	"time"

	// "github.com/garyburd/redigo/redis"
	"github.com/gomodule/redigo/redis"
)

type RedisCache struct {
	option *RedisConfig
	pool   *redis.Pool
}

func NewRedisCache(config *RedisConfig) *RedisCache {
	return &RedisCache{option: config}
}

// newRedisPool : 创建redis连接池
func (u *RedisCache) newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			// 1. 打开连接
			c, err := redis.Dial("tcp", u.option.Addr)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			// 2. 访问认证
			if _, err = c.Do("AUTH", u.option.Pass); err != nil {
				fmt.Println(err)
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func (u *RedisCache) GetPool() *redis.Pool {
	if u.pool == nil {
		u.pool = u.newRedisPool()
		u.pool.Get().Do("KEYS", "*")
	}
	return u.pool
}
