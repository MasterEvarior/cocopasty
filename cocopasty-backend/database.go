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
	}, redisConnection.Ping(context.Background()).Err()
}

func (d *database) CreateEntry(ctx context.Context, code string) error {
	log.Debug("Setting value in Redis")
	status := d.connection.Set(ctx, keyName, code, 0)

	return status.Err()
}

func (d *database) ReadEntry(ctx context.Context) (string, error) {
	log.Debug("Getting value from Redis")

	return d.connection.Get(ctx, keyName).Result()
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
	} else {
		log.Info("A Redis password was set")
	}

	return input
}
