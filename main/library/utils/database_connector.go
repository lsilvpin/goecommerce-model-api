package utils

import (
	"log"

	entities "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"
	"github.com/lsilvpin/goecommerce-model-api/main/library/tools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectToDatabaseAndMigrate() {
	connectionString := tools.GetEnv("DATABASE_CONNECTION_STRING")
	log.Println("Connection string: ", connectionString)
	err := error(nil)
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Failed to connect to database")
	}
	DB.AutoMigrate(&entities.Sample{})
}
