package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CodeSnippet struct {
	Code string `json:"Code"`
}

func main() {
	log.Print("Starting Cocopasty...")
	port := GetEnvVarWithDefault("COCOPASTY_PORT", ":8080")

	http.HandleFunc("/", handle)
	http.HandleFunc("/health", handleHealth)

	log.Print("Starting web server...")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGets(w, r)
	case http.MethodPost:
		handlePosts(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Length", "0")
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlePosts(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		log.Print("Invalid content type, returning 400")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newSnippet CodeSnippet

	err := json.NewDecoder(r.Body).Decode(&newSnippet)
	if err != nil {
		log.Printf("Could not decode JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = Persist(newSnippet.Code)
	if err != nil {
		log.Printf("Failed to save snippet in Redis: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleGets(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	value, err := Retrieve()

	if err != nil {
		log.Printf("Could not retrieve snippet: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Print("GET-Request successfull, returning 200")
	json.NewEncoder(w).Encode(CodeSnippet{value})
}
