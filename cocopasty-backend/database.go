package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var connection *redis.Client
var ctx context.Context

const keyName = "cocopasty-code-snippet"

func createConnection() {
	connection = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	ctx = context.Background()
}

func createEntry(code string) *redis.StatusCmd {
	if connection == nil {
		createConnection()
	}

	err := connection.Set(ctx, keyName, code, 0)
	if err != nil {
		fmt.Println(time.Now(), err)
	}

	return err
}
