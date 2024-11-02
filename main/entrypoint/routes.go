package entrypoint

import (
	"github.com/gin-gonic/gin"
	"github.com/lsilvpin/goecommerce-model-api/main/entrypoint/controllers"
)

func SetupApi() {
	r := gin.Default()
	r.GET("/samples", controllers.GetSamples)
	r.GET("/samples/:id", controllers.GetSampleById)
	r.POST("/samples", controllers.CreateSample)
	r.PUT("/samples/:id", controllers.UpdateSample)
	r.DELETE("/samples/:id", controllers.DeleteSample)
	r.Run()
}
