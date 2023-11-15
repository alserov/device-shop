package cache

import (
	"github.com/go-redis/redis"
	"log"
	"os"
	"time"
)

func Connect() (*redis.Client, error) {
	cl := redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_ADDR"),
		Password:    "",
		DB:          0,
		DialTimeout: 3 * time.Second,
		ReadTimeout: 3 * time.Second,
	})

	res, err := cl.Ping().Result()
	if err != nil {
		return nil, err
	}
	log.Println("REDIS PING ", res)

	log.Println("REDIS connected")
	return cl, nil
}
