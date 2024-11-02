package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	models "github.com/lsilvpin/goecommerce-model-api/main/library/domain/models"
	"github.com/lsilvpin/goecommerce-model-api/main/library/utils"
)

func GetSamples(c *gin.Context) {
	c.JSON(200, gin.H{
		"samples": utils.GenerateSampleList(),
		"retorno": models.ReturnModel{
			Trace:             "",
			Message:           "Lista de amostras retornada com sucesso",
			HttpStatusMessage: "OK",
			HttpStatusCode:    200,
		},
	})
}

func GetSampleById(c *gin.Context) {
	idFromInput := c.Param("id")
	sample := utils.GenerateSample()
	idFromInputAsInt, _ := strconv.Atoi(idFromInput)
	sample.Id = idFromInputAsInt
	c.JSON(200, gin.H{
		"sample": sample,
		"retorno": models.ReturnModel{
			Trace:             "",
			Message:           "Amostra retornada com sucesso",
			HttpStatusMessage: "OK",
			HttpStatusCode:    200,
		},
	})
}

func CreateSample(c *gin.Context) {
	createdSampleObj := map[string]interface{}{
		"id":   1,
		"name": "Sample Name",
	}
	c.JSON(201, gin.H{
		"id":   createdSampleObj["id"],
		"name": createdSampleObj["name"],
	})
}

func UpdateSample(c *gin.Context) {
	id := c.Param("id")
	updatedSampleObj := map[string]interface{}{
		"id":   id,
		"name": "Sample Name Updated",
	}
	c.JSON(200, gin.H{
		"id":   updatedSampleObj["id"],
		"name": updatedSampleObj["name"],
	})
}

func DeleteSample(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Sample " + id + " deleted successfully",
	})
}
