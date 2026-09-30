package cache

import (
	"context"
	"log"
	"time"

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

func Get(shortCode string) (string, bool) {
	ctx := context.Background()

	val, err := RDB.Get(ctx, shortCode).Result()
	if err == redis.Nil {
		//key doesnt exist
		return "", false
	}
	if err != nil {
		// redis down etc....
		log.Println("Redis Get error:", err)
		return "", false
	}

	return val, true
}

func Set(shortCode, longURL string) {
	ctx := context.Background()

	err := RDB.Set(ctx, shortCode, longURL, 24*time.Hour).Err()
	if err != nil {
		log.Println("redis set error :", err)

	}
}
