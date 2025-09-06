package main

import (
	"fmt"
	"log"
)

func formatActivity(activity *StravaActivity) string {
	distance := activity.Distance / 1000.0
	return fmt.Sprintf("#Strava\nNew Strava Activity!\nTitle: %s\nDistance: %.2f km", activity.Name, distance)
}

func (cfg *apiConfig) activityHandler(webhook_event *WebhookEvent) {
	// for now if its a update or delete do nothing
	switch webhook_event.AspectType {
	case "create":
		log.Println("Event was created, calling strava api")
		activity, err := cfg.getActivity(webhook_event)
		if err != nil {
			log.Printf("Could not get activity: %v", err)
			return
		}
		payload := MemosPayload{
			State:      "NORMAL",
			Content:    formatActivity(&activity),
			Visibility: "PROTECTED",
		}
		cfg.PostMemo(payload)

	case "update":
		log.Println("Event was updated")
		log.Println("Event was updated, calling strava api")
		activity, err := cfg.getActivity(webhook_event)
		if err != nil {
			log.Printf("Could not get activity: %v", err)
			return
		}
		payload := MemosPayload{
			State:      "NORMAL",
			Content:    activity.Name,
			Visibility: "PROTECTED",
		}
		cfg.PostMemo(payload)
	case "delete":
		log.Println("Event was delete")
	default:
		log.Println("Event not recognised doing nothing.")
	}

}
