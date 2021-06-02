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
	sub := rdb.Subscribe(context.TODO(), "queue")
	ch := sub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

}
