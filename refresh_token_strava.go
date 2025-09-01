package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (cfg *apiConfig) refreshStravaToken() (*StravaTokenResponse, error) {

	apiURL := "https://www.strava.com/api/v3/oauth/token"

	form := url.Values{}
	form.Set("client_id", cfg.ClientID)
	form.Set("client_secret", cfg.ClientSecret)
	form.Set("grant_type", "refresh_token")
	form.Set("refresh_token", cfg.RefreshToken)

	client := &http.Client{}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(form.Encode())))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to send request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response returned non-200 status code: %d", resp.StatusCode)
	}

	var tokenResponse StravaTokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal data: %w", err)
	}

	return &tokenResponse, nil
}
