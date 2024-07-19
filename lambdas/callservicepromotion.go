package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-endpoint:6379", // Update with your Redis endpoint
		Password: "",                    // no password set
		DB:       0,                     // use default DB
	})
}

func callAPI2(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func handleRequest2(w http.ResponseWriter, r *http.Request) {
	apiResponse, err := callAPI2("https://api2/data")
	if err != nil {
		http.Error(w, fmt.Sprintf("error calling API 2: %v", err), http.StatusInternalServerError)
		return
	}

	err = redisClient.Set(ctx, "api2_response", apiResponse, 0).Err()
	if err != nil {
		http.Error(w, fmt.Sprintf("error setting Redis key: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "API 2 Response: %s", apiResponse)
}
