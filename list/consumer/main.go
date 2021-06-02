package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	for true {
		//result, err := rdb.RPop(context.Background(), "queue").Result()
		result, err := rdb.BRPop(context.Background(), 0, "queue").Result()
		if err == redis.Nil {
			time.Sleep(2 * time.Second)
			continue
		} else if err != nil {
			log.Println(err)
		}
		fmt.Printf("consumer msg %v\n", result)

	}
}
