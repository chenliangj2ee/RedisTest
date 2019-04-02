package db

import (
	"time"
	"github.com/garyburd/redigo/redis"
)

var RedisPool *redis.Pool

const HOST = "127.0.0.1"
const PORT = "6379"

func init() {
	RedisPool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   20,
		IdleTimeout: 3 * 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", HOST+":"+PORT)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err

		},
	}

}
