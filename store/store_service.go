package store

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const cacheDurtion = 6 * time.Hour

func InitializeStore() *StorageService {
	redisIP := os.Getenv("REDIS_HOST")
	if redisIP == "" {
		redisIP = "localhost"
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisIP + ":6379",
		Password: "",
		DB:       0,
	})
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

/*
We want to be able to save the mapping between the originalUrl
and the generated shortUrl url
*/
func SaveURLMapping(shortURL, originalURL string) {
	err := storeService.redisClient.Set(ctx, shortURL, originalURL, cacheDurtion).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortURL, originalURL))
	}
}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/
func RetrieveInitiaURL(shortURL string) string {
	result, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialURL url | Error: %v - shortUrl: %s\n", err, shortURL))
	}
	return result
}
