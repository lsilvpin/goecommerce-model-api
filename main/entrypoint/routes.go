package entrypoint

import (
	"github.com/gin-gonic/gin"
	"github.com/lsilvpin/goecommerce-model-api/main/entrypoint/controllers"
)

func SetupApi() {
	r := gin.Default()
	r.GET("/info", controllers.GetInfo)
	r.Run()
}
