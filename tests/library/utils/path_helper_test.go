package tests

import (
	"testing"

	"github.com/lsilvpin/goecommerce-model-api/main/library/utils"
)

func TestShouldGetRootDirSuccessfully(t *testing.T) {
	// Arrange
	// Act
	rootDir, err := utils.GetRootDir()

	// Assert
	utils.AssertTrue(t, err == nil, "Error getting root dir")
	utils.AssertTrue(t, rootDir != "", "Root dir is empty")
}
