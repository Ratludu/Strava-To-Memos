package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const OAuthURL = "https://www.strava.com/api/v3/oauth/token"

func (cfg *apiConfig) handlerOk(w http.ResponseWriter, r *http.Request) {

	var event WebhookEvent

	log.Println("Received POST request on /strava-webhook")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&event)
	if err != nil {
		log.Printf("Could not decode request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	strSubID, err := strconv.Atoi(cfg.SubscriptionID)
	if err != nil {
		log.Printf("Could not convert subscription_id to string: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if strSubID != event.SubscriptionID {
		log.Print("Request did not have the same SubscriptionID")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if event.ObjectType != "activity" {
		log.Print("Webhook related to athelete data, ignoring")
		http.Error(w, "Not taking athelete data now", http.StatusBadRequest)
		return
	}

	go cfg.activityHandler(&event)

	log.Println("Responding with OK")
	w.WriteHeader(http.StatusOK)
}

func (cfg *apiConfig) getActivity(webhook_event *WebhookEvent) (StravaActivity, error) {

	activityURL := fmt.Sprintf("https://www.strava.com/api/v3/activities/%d", webhook_event.ObjectID)

	req, err := http.NewRequest("GET", activityURL, nil)
	if err != nil {
		return StravaActivity{}, fmt.Errorf("Error: Could not create request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return StravaActivity{}, fmt.Errorf("Error doing request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return StravaActivity{}, fmt.Errorf("Error: Could not read body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return StravaActivity{}, fmt.Errorf("Error: Could not get data, api failed with error code %d", resp.StatusCode)
	}

	var stravaActivity StravaActivity
	if err = json.Unmarshal(body, &stravaActivity); err != nil {
		return StravaActivity{}, fmt.Errorf("Error: Failed to unmarshal strava data: %v", err)
	}

	return stravaActivity, nil
}
