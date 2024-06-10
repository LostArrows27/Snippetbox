package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	value := os.Getenv(name)

	if value == "" {
		// logger.Error(errors.New("not found port"))
		log.Fatal("not found " + name)
	}

	return value
}

func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		// logger.Error(errors.New("error in loading .env"))
		log.Fatal("error in loading .env")

	}
}
