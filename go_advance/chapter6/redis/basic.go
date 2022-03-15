package main

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "120.24.250.251:6379",
		Password: "eeeecj_23_redis",
		DB:       0,
	})
	fmt.Println(rdb.Keys(context.Background(), "*"))
}
