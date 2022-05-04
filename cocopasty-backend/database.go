package main

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var connection *redis.Client
var ctx context.Context

const keyName = "cocopasty-code-snippet"

func createConnection() {
	setLogLevel()
	log.Debug("Creating connection to Redis...")
	connection = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     getAddress(),
		Password: getPassword(),
		DB:       0,
	})

	ctx = context.Background()
}

func createEntry(code string) *redis.StatusCmd {
	if connection == nil {
		createConnection()
	}

	log.Debug("Setting value in Redis...")
	err := connection.Set(ctx, keyName, code, 0)
	if err != nil {
		log.Error(err)
	}

	return err
}

func readEntry() (string, bool) {
	if connection == nil {
		createConnection()
	}

	log.Debug("Getting value from Redis...")
	stringValue, err := connection.Get(ctx, keyName).Result()

	if err != nil {
		log.Error(err)
		return stringValue, true
	}

	return stringValue, false
}

func getAddress() string {
	input := os.Getenv("REDIS_HOST")
	log.Info("REDIS_HOST: ", input)

	return input
}

func getPassword() string {
	input := os.Getenv("REDIS_PASSWORD")
	if input == "" {
		log.Info("No Redis password was set")
	}

	return input
}
