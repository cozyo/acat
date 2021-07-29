package config

import (
	"fmt"
	"github.com/acat/core/config/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

func NewRedis(settings *setting.RedisSettings) (*redis.Pool, error) {
	if settings == nil {
		return nil, fmt.Errorf("Redis setting is nil")
	}
	if settings.Host == "" {
		return nil, fmt.Errorf("Lack of redis setting Host")
	}
	if settings.Password == "" {
		return nil, fmt.Errorf("Lack of redis setting Password")
	}
	if settings.PoolNum <= 0 {
		return nil, fmt.Errorf("Wrong redis setting PoolNum config")
	}

	return &redis.Pool{
		MaxIdle: settings.PoolNum,
		MaxActive: settings.PoolNum,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", settings.Host)
			if err != nil {
				return nil, err
			}
			if settings.Password != "" {
				if _, err := c.Do("AUTH", settings.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if settings.DataBase != 0 {
				if _, err := c.Do("SELECT", settings.DataBase); err != nil {
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}