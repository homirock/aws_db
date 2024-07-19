// handler.go
package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-endpoint:6379", // Update with your Redis endpoint
		Password: "",                    // no password set
		DB:       0,                     // use default DB
	})
}

func handleRequest3() (string, error) {
	api1Response, err := redisClient.Get(ctx, "api1_response").Result()
	if err != nil {
		return "", fmt.Errorf("error getting API 1 response from Redis: %v", err)
	}

	api2Response, err := redisClient.Get(ctx, "api2_response").Result()
	if err != nil {
		return "", fmt.Errorf("error getting API 2 response from Redis: %v", err)
	}

	combinedResponse := fmt.Sprintf("API 1 Response: %s\nAPI 2 Response: %s", api1Response, api2Response)
	fmt.Println(combinedResponse)

	return combinedResponse, nil
}
