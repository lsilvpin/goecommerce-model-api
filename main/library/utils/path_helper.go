package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func GetRootDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		envPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(envPath); err == nil {
			return currentDir, nil
		}
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", errors.New("root directory not found")
		}
		currentDir = parentDir
	}
}
