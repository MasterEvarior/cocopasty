package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func setLogLevel() {
	input := os.Getenv("LOG_LEVEL")

	parsedLevel, err := log.ParseLevel(input)

	if err == nil {
		log.SetLevel(parsedLevel)
	}

	log.Info("Log level to ", log.GetLevel())
}
