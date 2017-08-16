package main

import (
	"encoding/json"
	"net/smtp"
	"os"
	"strings"

	"github.com/thcdrt/funk-go/event"
)

type Data struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Text    string   `json:"text"`
}

func Send(event event.Event) (string, error) {
	host := os.Getenv("MAIL_SMTP_HOST")
	port := os.Getenv("MAIL_SMTP_PORT")
	username := os.Getenv("MAIL_SMTP_USERNAME")
	password := os.Getenv("MAIL_SMTP_PASSWORD")

	var data Data
	err := json.Unmarshal([]byte(event.Data), &data)
	if err != nil {
		return "", err
	}

	msg := []byte("To: " + strings.Join(data.To[:], " ") + "\r\n" +
		"Subject: " + data.Subject + "\r\n" +
		"\r\n" +
		data.Text + "\r\n",
	)

	auth := smtp.PlainAuth("", username, password, host)
	err = smtp.SendMail(host+":"+port, auth, username, data.To, msg)
	if err != nil {
		return "", err
	}

	return "Mail sent", nil
}
