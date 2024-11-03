package repositories

import (
	"errors"

	entities "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"
	"github.com/lsilvpin/goecommerce-model-api/main/library/utils"
)

func CreateSample(sample entities.Sample) {
	utils.DB.Create(&sample)
}

func ReadAllSamples() []entities.Sample {
	samples := []entities.Sample{}
	utils.DB.Find(&samples)
	return samples
}

func ReadSampleById(id uint64) (entities.Sample, error) {
	sample := entities.Sample{}
	utils.DB.First(&sample, id)
	var err error
	if sample.ID == 0 {
		err = errors.New("sample not found")
	}
	return sample, err
}

func UpdateSample(id uint64, sample entities.Sample) (entities.Sample, error) {
	sampleFromDb, err := ReadSampleById(id)
	if err == nil {
		sampleFromDb.Name = sample.Name
		sampleFromDb.Age = sample.Age
		sampleFromDb.Size = sample.Size
		sampleFromDb.IsVisible = sample.IsVisible
		utils.DB.Save(&sampleFromDb)
	}
	return sampleFromDb, err
}

func DeleteSampleById(id uint64) error {
	sample, err := ReadSampleById(id)
	if err == nil {
		utils.DB.Delete(&sample)
	}
	return err
}
