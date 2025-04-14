package main

import (
	"os"
)

var inMemoryStorage string = ""

func Persist(content string) error {
	if !toDisk() {
		inMemoryStorage = content
		return nil
	}

	err := os.WriteFile(location(), []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func Retrieve() (string, error) {
	if !toDisk() {
		return inMemoryStorage, nil
	}

	b, err := os.ReadFile(location())
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func toDisk() bool {
	return IsEnabled("COCOPASTY_PERSIST_TO_DISK")
}

func location() string {
	return GetEnvVarWithDefault("COCOPASTY_DATA_LOCATION", "/data/value")
}
