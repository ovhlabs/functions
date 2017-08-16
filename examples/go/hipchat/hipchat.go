package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/thcdrt/funk-go/event"
)

type Data struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	MessageFormat string `json:"message_format"`
}

func Notify(event event.Event) (string, error) {
	roomID := os.Getenv("HIPCHAT_ROOM_ID")
	authToken := os.Getenv("HIPCHAT_AUTH_TOKEN")
	url := "https://api.hipchat.com/v2/room/" + roomID + "/notification?auth_token=" + authToken

	// parse the event.Data string
	var data Data
	err := json.Unmarshal([]byte(event.Data), &data)
	if err != nil {
		return "", err
	}

	// prepare the request body
	data.MessageFormat = "html"
	if data.Message == "" {
		return "", errors.New("message not found")
	}
	if data.Color == "" {
		data.Color = "green"
	}
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// send the request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// check the response status
	if resp.StatusCode != 204 {
		return "", errors.New("failed to send " + data.Message)
	}

	return "OK", err
}
