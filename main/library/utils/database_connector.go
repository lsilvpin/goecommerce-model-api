package utils

import (
	"log"
	"path/filepath"

	"github.com/glebarez/sqlite"
	entities "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectToDatabaseAndMigrate() error {
	environment := GetEnv("ENVIRONMENT")
	log.Println("Environment: ", environment)
	connectionString := GetEnv("DATABASE_CONNECTION_STRING")
	log.Println("Connection string: ", connectionString)
	var err error
	if environment == "dev" {
		rootDir, err := GetRootDir()
		if err != nil {
			log.Panic("Root directory not found")
			return err
		}
		sqliteDbPath := filepath.Join(rootDir, ".tmp", "db.sqlite")
		DB, err = gorm.Open(sqlite.Open(sqliteDbPath), &gorm.Config{})
		if err != nil {
			log.Panic("Failed to connect to sqlite")
			return err
		}
		log.Println("Connected to sqlite")
	} else {
		DB, err = gorm.Open(postgres.Open(connectionString))
		if err != nil {
			log.Panic("Failed to connect to postgres")
			return err
		}
		log.Println("Connected to postgres")
	}
	DB.AutoMigrate(&entities.Sample{})
	return nil
}
