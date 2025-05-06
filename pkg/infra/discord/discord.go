package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Discord struct {
	webhookURL string
}

func NewDiscord(webhookURL string) Discord {
	return Discord{
		webhookURL: webhookURL,
	}
}

func (d Discord) SendMessage(message string) error {
	payload := map[string]string{"content": message}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(d.webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("failed to send: %s", string(body))
	}
	return nil
}
