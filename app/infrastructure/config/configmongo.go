package config

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(ctx context.Context) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(getEnvVariable("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	logrus.Info("Connected to MongoDB")

	return client
}

// getting database collections
func GetFilialCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(getEnvVariable("FILIAL_DATABASE")).Collection(getEnvVariable("FILIAL_COLLECTION"))
}
