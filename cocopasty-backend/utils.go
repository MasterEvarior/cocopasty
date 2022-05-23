package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func setLogLevel() {
	defer log.Info("Log level to ", log.GetLevel())
	input := os.Getenv("LOG_LEVEL")

	parsedLevel, err := log.ParseLevel(input)

	if err != nil {
		log.Error(err)
		return
	}

	log.SetLevel(parsedLevel)
}
