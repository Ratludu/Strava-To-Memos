package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const postMemoUrl = "memos"

type MemosPayload struct {
	State      string `json:"state"`
	Content    string `json:"content"`
	Visibility string `json:"visibility"`
}

func (cfg *apiConfig) PostMemo(payload MemosPayload) error {

	urlPath, err := cfg.ExtendUrl(postMemoUrl)
	if err != nil {
		return err
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", urlPath.String(), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.MemosApiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request.")
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Memo not created, status code: %d", resp.StatusCode)
	}

	return nil

}
