package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/runabove/functions/go-sdk/event"
)

func Pub(event event.Event) (string, error) {
	topic := os.Getenv("KAFKA_TOPIC")
	host := os.Getenv("KAFKA_HOST")
	user := os.Getenv("KAFKA_USER")
	password := os.Getenv("KAFKA_PASSWORD")

	if event.Data == "" {
		return "", errors.New("data not found")
	}

	var config = sarama.NewConfig()
	config.Net.TLS.Enable = true
	config.Net.SASL.Enable = true
	config.Net.SASL.User = user
	config.Net.SASL.Password = password
	config.ClientID = user
	config.Producer.Return.Successes = true

	var err error
	producer, err := sarama.NewSyncProducer([]string{host}, config)
	if err != nil {
		return "", err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(event.Data)}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Message '%s' sent to partition %d at offset %d\n", event.Data, partition, offset), nil
}
