package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	MemosApiKey string
	MemosURL    string
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Couldn't find a .env file, assuming default has been set")
	}

	apiCfg := apiConfig{
		MemosApiKey: os.Getenv("MEMOS_API"),
		MemosURL:    os.Getenv("MEMOS_URL"),
	}

	testMemo := MemosPayload{
		Content:    "testing",
		State:      "NORMAL",
		Visibility: "PROTECTED",
	}

	err = apiCfg.PostMemo(testMemo)
	if err != nil {
		fmt.Println("Couldn't post memo:", err)
		return
	}

}
