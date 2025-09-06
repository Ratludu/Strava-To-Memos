package main

import (
	"time"
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

type StravaData struct {
	RecentRunTotals struct {
		Count            int     `json:"count"`
		Distance         float64 `json:"distance"`
		MovingTime       int     `json:"moving_time"`
		ElapsedTime      int     `json:"elapsed_time"`
		ElevationGain    float64 `json:"elevation_gain"`
		AchievementCount int     `json:"achievement_count"`
	} `json:"recent_run_totals"`
	AllRunTotals struct {
		Count         int     `json:"count"`
		Distance      float64 `json:"distance"`
		MovingTime    int     `json:"moving_time"`
		ElapsedTime   int     `json:"elapsed_time"`
		ElevationGain float64 `json:"elevation_gain"`
	} `json:"all_run_totals"`
	YtdRunTotals struct {
		Count         int     `json:"count"`
		Distance      int     `json:"distance"`
		MovingTime    float64 `json:"moving_time"`
		ElapsedTime   float64 `json:"elapsed_time"`
		ElevationGain float64 `json:"elevation_gain"`
	} `json:"ytd_run_totals"`
}

type StravaTokenResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresAt    int64  `json:"expires_at"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Athlete      struct {
		ID int64 `json:"id"`
	} `json:"athlete"`
}

type StravaActivity struct {
	ID            int64  `json:"id"`
	ResourceState int    `json:"resource_state"`
	ExternalID    string `json:"external_id"`
	UploadID      int64  `json:"upload_id"`
	Athlete       struct {
		ID            int `json:"id"`
		ResourceState int `json:"resource_state"`
	} `json:"athlete"`
	Name               string    `json:"name"`
	Distance           float64   `json:"distance"`
	MovingTime         int       `json:"moving_time"`
	ElapsedTime        int       `json:"elapsed_time"`
	TotalElevationGain float64   `json:"total_elevation_gain"`
	Type               string    `json:"type"`
	SportType          string    `json:"sport_type"`
	StartDate          time.Time `json:"start_date"`
	StartDateLocal     time.Time `json:"start_date_local"`
	Timezone           string    `json:"timezone"`
	UtcOffset          float64   `json:"utc_offset"`
	StartLatlng        []float64 `json:"start_latlng"`
	EndLatlng          []float64 `json:"end_latlng"`
	Calories           float64   `json:"calories"`
}
