package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 1")
}
