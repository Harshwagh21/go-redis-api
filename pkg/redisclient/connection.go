package redisclient

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Init() {
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		log.Fatal("REDIS_URL environment variable is not set")
	}

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatal("Failed to parse Redis URL: ", err)
	}

	Client = redis.NewClient(opt)

	pong, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis: ", err)
	}

	fmt.Printf("Redis is connected: %s\n", pong)
}
