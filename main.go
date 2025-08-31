package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	MemosApiKey       string
	MemosURL          string
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
		StravaVerifyToken: os.Getenv("VERIFY_TOKEN"),
	}

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
