package main

import (
	"github.com/go-redis/redis/v8"
)

func createConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func createEntry(code string, language string) {
}
