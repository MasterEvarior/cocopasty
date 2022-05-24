package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Debugf("Incoming request: \n Method: '%s' \n URI: '%s'", r.Method, r.RequestURI)
			next.ServeHTTP(w, r)
		},
	)
}
