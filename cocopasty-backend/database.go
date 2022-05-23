package main

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type database struct {
	connection *redis.Client
}

const keyName = "cocopasty-code-snippet"

func CreateDatabaseClient() (*database, error) {
	setLogLevel()
	log.Debug("Creating connection to Redis...")

	redisConnection := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     getAddress(),
		Password: getPassword(),
		DB:       0,
	})

	return &database{
		connection: redisConnection,
	}, nil
}

func (d *database) CreateEntry(ctx context.Context, code string) *redis.StatusCmd {
	log.Debug("Setting value in Redis...")
	err := d.connection.Set(ctx, keyName, code, 0)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (d *database) ReadEntry(ctx context.Context) (string, bool) {
	log.Debug("Getting value from Redis...")
	stringValue, err := d.connection.Get(ctx, keyName).Result()

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
