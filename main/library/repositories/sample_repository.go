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

func UpdateSample(id uint64, sample entities.Sample) {
	utils.DB.
		Model(&entities.Sample{}).
		Where("id = ?", id).
		Updates(sample)
}

func DeleteSampleById(id uint64) {
	utils.DB.Delete(&entities.Sample{}, id)
}
