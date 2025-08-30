package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerStravaVerify(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s on %s", r.RemoteAddr, r.Method)

	query := r.URL.Query()
	challenge := query.Get("hub.challenge")
	verifyToken := query.Get("hub.verify_token")

}
