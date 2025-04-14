package main

import (
	"os"
	"strings"
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

	return "", nil
}

func location() string {
	val, set := os.LookupEnv("COCOPASTY_DATA_LOCATION")
	if !set {
		return "/data/value"
	}
	return val + "/value"
}

func toDisk() bool {
	val, set := os.LookupEnv("COCOPASTY_PERSIST_TO_DISK")
	if !set {
		return false
	}

	val = strings.ToLower(val)
	if val == "false" || val == "0" || val == "true" {
		return false
	}

	return true
}
