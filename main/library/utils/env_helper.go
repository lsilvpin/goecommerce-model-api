package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	rootDir, err := GetRootDir()
	if err != nil {
		log.Panic("Error getting root dir: ", err)
	}
	envFilePath := filepath.Join(rootDir, ".env")
	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Panic("Error loading .env file: ", err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func SetEnv(key string, value string) {
	os.Setenv(key, value)
}
