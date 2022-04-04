package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type CodeSnippet struct {
	Code     string `json:Code`
	Language string `json:Language`
}

func main() {
	//Initialize router
	router := mux.NewRouter()

	router.HandleFunc("/", handleGets).Methods("GET")
	router.HandleFunc("/", handlePosts).Methods("POST")

	corsHandler := cors.Default().Handler(router)

	//Start server
	err := http.ListenAndServe(":8080", corsHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func handlePosts(w http.ResponseWriter, r *http.Request) {

}

func handleGets(w http.ResponseWriter, r *http.Request) {
	snippet := CodeSnippet{"Console.log(12)", "js"}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(snippet)
}
