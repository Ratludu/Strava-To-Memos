package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerStravaVerify(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s on %s", r.RemoteAddr, r.Method)

	query := r.URL.Query()
	challenge := query.Get("hub.challenge")
	verifyToken := query.Get("hub.verify_token")

	if verifyToken != cfg.StravaVerifyToken {
		http.Error(w, "Verification failed", http.StatusUnauthorized)
		log.Println("Webhook verification failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	params := map[string]string{"hub.challenge": challenge}
	json.NewEncoder(w).Encode(params)
	log.Println("Webhook verification successful")
}
