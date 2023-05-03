package config

import (
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func ConnectKafka() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{getEnvVariable("BROKER_ADDRESS")},
		Topic:   getEnvVariable("TOPIC"),
		GroupID: getEnvVariable("GROUP_ID"),
		Logger:  log.New(os.Stdout, "kafka reader: ", 0),
	})
}
