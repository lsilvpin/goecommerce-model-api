package utils

import (
	"log"

	entities "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectToDatabaseAndMigrate() {
	connectionString := "host=localhost user=postgres password=postgres dbname=goecommerce port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	err := error(nil)
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Failed to connect to database")
	}
	DB.AutoMigrate(&entities.Sample{})
}
