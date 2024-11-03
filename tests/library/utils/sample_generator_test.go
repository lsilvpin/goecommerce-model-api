package tests

import (
	"testing"

	entities "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"
	"github.com/lsilvpin/goecommerce-model-api/main/library/utils"
)

func TestShouldGenerateSampleSuccessfully(t *testing.T) {
	// Arrange
	expectedSample := entities.Sample{
		Name:      "Sample Name",
		Age:       20,
		Size:      1.75,
		IsVisible: true,
	}

	// Act
	sampleGenerated := utils.GenerateSample()

	// Assert
	utils.AssertTrue(t, expectedSample.Name == sampleGenerated.Name, "Name is different")
	utils.AssertTrue(t, expectedSample.Age == sampleGenerated.Age, "Age is different")
	utils.AssertTrue(t, expectedSample.Size == sampleGenerated.Size, "Size is different")
	utils.AssertTrue(t, expectedSample.IsVisible == sampleGenerated.IsVisible, "IsVisible is different")
}

func TestShouldGenerateManySamplesSuccefully(t *testing.T) {
	// Arrange
	expectedSampleList := []entities.Sample{
		{
			Name:      "Sample Name",
			Age:       20,
			Size:      1.75,
			IsVisible: true,
		},
		{
			Name:      "Sample Name",
			Age:       20,
			Size:      1.75,
			IsVisible: true,
		},
	}

	// Act
	sampleListGenerated := utils.GenerateSampleList(2)

	// Assert
	utils.AssertTrue(t, len(sampleListGenerated) == 2, "Length is different")
	utils.AssertTrue(t, expectedSampleList[0].Name == sampleListGenerated[0].Name, "Name is different")
	utils.AssertTrue(t, expectedSampleList[0].Age == sampleListGenerated[0].Age, "Age is different")
	utils.AssertTrue(t, expectedSampleList[0].Size == sampleListGenerated[0].Size, "Size is different")
	utils.AssertTrue(t, expectedSampleList[0].IsVisible == sampleListGenerated[0].IsVisible, "IsVisible is different")
	utils.AssertTrue(t, expectedSampleList[1].Name == sampleListGenerated[1].Name, "Name is different")
	utils.AssertTrue(t, expectedSampleList[1].Age == sampleListGenerated[1].Age, "Age is different")
	utils.AssertTrue(t, expectedSampleList[1].Size == sampleListGenerated[1].Size, "Size is different")
	utils.AssertTrue(t, expectedSampleList[1].IsVisible == sampleListGenerated[1].IsVisible, "IsVisible is different")
}
