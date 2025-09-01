package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	// Memos
	MemosApiKey string
	MemosURL    string
	// Strava
	ClientID          string
	ClientSecret      string
	RefreshToken      string
	AccessToken       string
	StravaVerifyToken string
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Couldn't find a .env file, assuming default has been set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT envrionment variable was not set")
	}

	apiCfg := apiConfig{
		MemosApiKey:       os.Getenv("MEMOS_API"),
		MemosURL:          os.Getenv("MEMOS_URL"),
		ClientID:          os.Getenv("CLIENT_ID"),
		ClientSecret:      os.Getenv("CLIENT_SECRET"),
		RefreshToken:      os.Getenv("REFRESH_TOKEN"),
		StravaVerifyToken: os.Getenv("VERIFY_TOKEN"),
	}

	// get strava access token
	stravaResponse, err := apiCfg.refreshStravaToken()
	if err != nil {
		log.Fatal("Could not get strava access token.")
	}

	apiCfg.AccessToken = stravaResponse.AccessToken

	mux := http.NewServeMux()
	mux.HandleFunc("POST /strava-webhook", apiCfg.handlerOk)
	mux.HandleFunc("GET /strava-webhook", apiCfg.handlerStravaVerify)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving and listening on port %s", port)
	log.Fatal(srv.ListenAndServe())
}
