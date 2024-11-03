package repositories

import (
	entities "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"
	"github.com/lsilvpin/goecommerce-model-api/main/library/utils"
)

func CreateSample(sample entities.Sample) {
	utils.DB.Create(&sample)
}

func ReadSampleById(id uint64) entities.Sample {
	sample := entities.Sample{}
	utils.DB.First(&sample, id)
	return sample
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
