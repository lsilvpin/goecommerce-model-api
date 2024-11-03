package tests

import (
	"testing"

	"github.com/lsilvpin/goecommerce-model-api/main/library/utils"
)

func TestShouldGetEnvironmentVariableSuccessfully(t *testing.T) {
	// Arrange
	key := "TEST_ENV_KEY"
	value := "TEST_ENV_VALUE"
	utils.SetEnv(key, value)

	// Act
	envValue := utils.GetEnv(key)

	// Assert
	utils.AssertTrue(t, envValue == value, "Error getting environment variable")
}

func TestShouldProveThatEnvVarIsMorePriorityThanDotEnv(t *testing.T) {
	// Arrange
	key := "DATABASE_CONNECTION_STRING"
	value := "TEST_ENV_VALUE"
	utils.SetEnv(key, value)
	utils.LoadDotEnv()

	// Act
	envValue := utils.GetEnv(key)

	// Assert
	utils.AssertTrue(t, envValue == value, "Error getting environment variable")
}

func TestShouldProveThatEnvVarCanBeGetFromDotEnv(t *testing.T) {
	// Arrange
	key := "ENVIRONMENT"
	expectedValue := "dev"
	utils.LoadDotEnv()

	// Act
	envValue := utils.GetEnv(key)

	// Assert
	utils.AssertTrue(t, envValue == expectedValue, "Error getting environment variable from DotEnvFile")
}
