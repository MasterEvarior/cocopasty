package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type CodeSnippet struct {
	Code string `json:Code`
}

var snippet CodeSnippet

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
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newSnippet CodeSnippet

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newSnippet)

	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	snippet = newSnippet
	createEntry(newSnippet.Code)
	w.WriteHeader(http.StatusCreated)
}

func handleGets(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(snippet)
}
