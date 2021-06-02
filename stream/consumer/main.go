package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	result, err := rdb.XRead(context.TODO(), &redis.XReadArgs{
		Streams: []string{"queue1", "0"},
		Count:   0,
		Block:   0,
	}).Result()
	if err != nil {
		panic(err)
	}
	for r := range result {
		for m := range result[r].Messages {
			fmt.Println(result[r].Messages[m].Values)
		}

	}

}
