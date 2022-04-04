package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type CodeSnippet struct {
	Code     string `json:Code`
	Language string `json:Language`
}

func main() {
	http.HandleFunc("/", handleRequests)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func handleRequests(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		snippet := CodeSnippet{"Console.log(12)", "js"}
		json.NewEncoder(w).Encode(snippet)
	case "POST":
		//TODO
	default:
		io.WriteString(w, "Sorry, only GET and POST methods are supported.")
	}
}
