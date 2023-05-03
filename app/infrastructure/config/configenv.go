package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func getEnvVariable(key string) string {
	value := os.Getenv(key)
	if value != "" {
		logrus.Info("Environment variable value loaded for key: " + key)
		return value
	}

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Info(".env file not found. Using environment variables")
	}
	return os.Getenv(key)
}
