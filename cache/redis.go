package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := RDB.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("Could not connect to Redis:", err)
	}

	log.Println("Connected to Redis")
}
