package utils

import (
	"log"

	"github.com/glebarez/sqlite"
	entities "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"
	"github.com/lsilvpin/goecommerce-model-api/main/library/tools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectToDatabaseAndMigrate() {
	environment := tools.GetEnv("ENVIRONMENT")
	connectionString := tools.GetEnv("DATABASE_CONNECTION_STRING")
	log.Println("Connection string: ", connectionString)
	err := error(nil)
	if environment == "dev" {
		DB, err = gorm.Open(sqlite.Open(".tmp/db.sqlite"), &gorm.Config{})
		log.Println("Connected to sqlite")
	} else {
		DB, err = gorm.Open(postgres.Open(connectionString))
		log.Println("Connected to postgres")
	}
	if err != nil {
		log.Panic("Failed to connect to database")
	}
	DB.AutoMigrate(&entities.Sample{})
}
