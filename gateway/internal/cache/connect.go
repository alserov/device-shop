package cache

import (
	"github.com/go-redis/redis"
	"time"
)

func MustConnect(addr string) *redis.Client {
	cl := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    "",
		DB:          0,
		DialTimeout: 3 * time.Second,
		ReadTimeout: 3 * time.Second,
	})

	_, err := cl.Ping().Result()
	if err != nil {
		panic("failed to ping redis: " + err.Error())
	}

	return cl
}
