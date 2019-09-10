package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

func New() (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})

	pong, err := rds.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(pong)

	return rds, nil
}
