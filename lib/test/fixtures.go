package test

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

func MockRedis() *redis.Client {
	mr, _ := miniredis.Run()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	return client
}
