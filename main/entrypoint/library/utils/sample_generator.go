package utils

func GenerateSample() map[string]interface{} {
	samplesMap := map[string]interface{}{
		"id":         1,
		"name":       "Sample Name",
		"age":        20,
		"size":       1.75,
		"is_visible": true,
		"created_at": "2021-01-01T00:00:00Z",
		"updated_at": "2021-01-01T00:00:00Z",
	}
	return samplesMap
}

func GenerateSampleList() []map[string]interface{} {
	sampleList := []map[string]interface{}{}
	for i := 0; i < 10; i++ {
		sample := GenerateSample()
		sampleList = append(sampleList, sample)
	}
	return sampleList
}
