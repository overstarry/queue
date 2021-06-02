package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	data := []string{`{"name": "jinzhu11221212321", "age": 18, "tags": ["tag211", "tag2"], "orgs": {"orga": "orga"}`, `{"name": "asd", "age": 18, "tags": ["tag211", "tag2"], "orgs": {"orga": "orga"}`}
	for i := range data {
		err := rdb.Publish(context.TODO(), "queue", data[i]).Err()
		if err != nil {
			panic(err)
		}
	}

}
