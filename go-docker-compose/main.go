package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "my-redis-container:6379",
		Password: "",
		DB:       0,
	})

	for _ = range time.Tick(time.Second) {
		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("key", val)
	}
}
