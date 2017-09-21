package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ovhlabs/functions/go-sdk/event"
)

func Pub(event event.Event) (string, error) {
	topic := os.Getenv("KAFKA_TOPIC")
	host := os.Getenv("KAFKA_HOST")
	user := os.Getenv("KAFKA_USER")
	password := os.Getenv("KAFKA_PASSWORD")

	if event.Data == "" {
		return "", errors.New("data not found")
	}

	// prepares the request
	client := &http.Client{}
	url := host + "/topic/" + topic + "?format=raw"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(event.Data)))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(user, password)

	// sends the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}
