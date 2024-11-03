package main

import (
	"github.com/lsilvpin/goecommerce-model-api/main/entrypoint"
	"github.com/lsilvpin/goecommerce-model-api/main/library/utils"
)

func main() {
	utils.LoadDotEnv()
	utils.ConnectToDatabaseAndMigrate()
	entrypoint.SetupApi()
}
