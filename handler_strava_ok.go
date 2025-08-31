package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type WebhookEvent struct {
	AspectType     string `json:"aspect_type"`
	EventTime      int    `json:"event_time"`
	ObjectID       int    `json:"object_id"`
	ObjectType     string `json:"object_type"`
	OwnerID        int    `json:"owner_id"`
	SubscriptionID int    `json:"subscription_id"`
	Updates        struct {
		Title string `json:"title"`
	} `json:"updates"`
}

func (cfg apiConfig) handlerOk(w http.ResponseWriter, r *http.Request) {

	var event WebhookEvent

	log.Println("Received POST request on /strava-webhook")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&event)
	if err != nil {
		log.Printf("Could not decode request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	payload := MemosPayload{
		State:      "NORMAL",
		Content:    event.ObjectType,
		Visibility: "PROTECTED",
	}

	go cfg.PostMemo(payload)

	log.Println("Responding with OK")
	w.WriteHeader(http.StatusOK)
}
