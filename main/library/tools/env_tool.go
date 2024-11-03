package tools

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func SetEnv(key string, value string) {
	os.Setenv(key, value)
}
