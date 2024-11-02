package utils

import domain "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"

func GenerateSample() domain.Sample {
	sampleObj := domain.Sample{
		Id:        1,
		Name:      "Sample Name",
		Age:       20,
		Size:      1.75,
		IsVisible: true,
		CreatedAt: "2021-01-01T00:00:00Z",
		UpdatedAt: "2021-01-01T00:00:00Z",
	}
	return sampleObj
}

func GenerateSampleList() []domain.Sample {
	sampleList := []domain.Sample{}
	for i := 0; i < 10; i++ {
		sample := GenerateSample()
		sampleList = append(sampleList, sample)
	}
	return sampleList
}
