package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "139.198.175.236:6379",
		Password: "chengyu520",
		DB:       0,
		PoolSize: 100,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func main() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		return
	}

	// val, err := rdb.Set(ctx, "name", "ranen12", 0).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)
}
