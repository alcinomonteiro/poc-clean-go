package config

import (
	"context"
	"time"

	"github.com/heptiolabs/healthcheck"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

func Health(m *mongo.Client, r *kafka.Reader) healthcheck.Handler {
	health := healthcheck.NewHandler()
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))
	health.AddReadinessCheck("mongodb", mongoCheck(m))
	health.AddReadinessCheck("kafka", kafkaCheck(r, &kafkaError{}))
	return health
}

func mongoCheck(m *mongo.Client) healthcheck.Check {
	return func() error {
		err := m.Ping(context.TODO(), nil)
		if err != nil {
			return err
		}
		return nil
	}
}

func kafkaCheck(r *kafka.Reader, kerr *kafkaError) healthcheck.Check {
	return func() error {
		erros := 0
		for i := 0; i < 10; i++ {
			erros += int(r.Stats().Errors)
			time.Sleep(time.Second / 2)
		}
		if erros > 0 {
			return kerr
		}
		return nil
	}
}

type kafkaError struct{}

func (kerr *kafkaError) Error() string {
	return "Unable to establish connection to consumer"
}
