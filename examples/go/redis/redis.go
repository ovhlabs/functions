package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/runabove/functions/go-sdk/event"
	redis "gopkg.in/redis.v5"
)

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Set(event event.Event) (string, error) {
	var data Data
	err := json.Unmarshal([]byte(event.Data), &data)
	if err != nil {
		return "", err
	}

	result, err := NewRedisClient().Set(data.Key, data.Value, 0).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}

func Get(event event.Event) (string, error) {
	var data Data
	err := json.Unmarshal([]byte(event.Data), &data)
	if err != nil {
		return "", err
	}

	value, err := NewRedisClient().Get(data.Key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func NewRedisClient() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")
	db, _ := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 0)

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       int(db),
	})

	return client
}
