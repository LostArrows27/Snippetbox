package env

import (
	"errors"
	"os"

	"github.com/LostArrows27/snippetbox/pkg/logger"
	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	portString := os.Getenv(name)

	if portString == "" {
		logger.Error(errors.New("not found port"))
	}

	return portString
}

func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		logger.Error(errors.New("error in loading .env"))
	}
}
