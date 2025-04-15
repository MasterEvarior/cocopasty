package main

import (
	"os"
)

const location string = "./data/value"

func Persist(content string) error {
	f, err := os.Create(location)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(content))
	if err != nil {
		return err
	}
	return nil
}

func Retrieve() (string, error) {
	b, err := os.ReadFile(location)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
