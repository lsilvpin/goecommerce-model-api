package utils

import domain "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"

func GenerateSample() domain.Sample {
	sampleObj := domain.Sample{
		Name:      "Sample Name",
		Age:       20,
		Size:      1.75,
		IsVisible: true,
	}
	return sampleObj
}

func GenerateSampleList(numberOfSamples int) []domain.Sample {
	sampleList := []domain.Sample{}
	for i := 0; i < numberOfSamples; i++ {
		sample := GenerateSample()
		sampleList = append(sampleList, sample)
	}
	return sampleList
}
