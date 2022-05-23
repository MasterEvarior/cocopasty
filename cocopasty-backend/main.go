package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

type CodeSnippet struct {
	Code string `json:"Code"`
}

var databaseClient *database

func main() {
	setLogLevel()
	//Initialize router
	log.Info("Starting Cocopasty...")
	router := mux.NewRouter()

	router.HandleFunc("/", handleGets).Methods("GET")
	router.HandleFunc("/", handlePosts).Methods("POST")

	corsHandler := cors.Default().Handler(router)

	//Create database connection
	var err error
	databaseClient, err = CreateDatabaseClient()

	if err != nil {
		panic(err)
	}

	//Start server
	log.Info("Starting web server...")
	err = http.ListenAndServe(":8080", corsHandler)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Cocopasty is started and ready!")
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	log.Debug("Received POST-Request")

	ctx := r.Context()

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		log.Debug("Invalid content type, returning 400")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newSnippet CodeSnippet

	err := json.NewDecoder(r.Body).Decode(&newSnippet)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = databaseClient.CreateEntry(ctx, newSnippet.Code)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Debug("GET-Request successfull, returning 200")
}

func handleGets(w http.ResponseWriter, r *http.Request) {
	log.Debug("Received GET-Request")

	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	value, err := databaseClient.ReadEntry(ctx)

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Debug("GET-Request successfull, returning 200")
	json.NewEncoder(w).Encode(CodeSnippet{value})
}
