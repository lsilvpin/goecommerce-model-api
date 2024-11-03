package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	entities "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"
	models "github.com/lsilvpin/goecommerce-model-api/main/library/domain/models"
	"github.com/lsilvpin/goecommerce-model-api/main/library/repositories"
)

func GetSamples(c *gin.Context) {
	samples := repositories.ReadAllSamples()
	c.JSON(200, gin.H{
		"samples": samples,
		"retorno": models.ReturnModel{
			Trace:             "",
			Message:           "Amostras retornadas com sucesso",
			HttpStatusMessage: "OK",
			HttpStatusCode:    200,
		},
	})
}

func GetSampleById(c *gin.Context) {
	var idFromInput uint64
	idFromInput, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"retorno": models.ReturnModel{
				Trace:             "",
				Message:           "Erro ao buscar amostra: " + err.Error(),
				HttpStatusMessage: "Internal Server Error",
				HttpStatusCode:    500,
			},
		})
		return
	}
	sample := repositories.ReadSampleById(idFromInput)
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
	var sample entities.Sample
	if err := c.ShouldBindJSON(&sample); err != nil {
		c.JSON(500, gin.H{
			"retorno": models.ReturnModel{
				Trace:             "",
				Message:           "Erro ao criar amostra: " + err.Error(),
				HttpStatusMessage: "Internal Server Error",
				HttpStatusCode:    500,
			},
		})
		return
	}
	repositories.CreateSample(sample)
	c.JSON(http.StatusOK, gin.H{
		"createdSample": sample,
		"retorno": models.ReturnModel{
			Trace:             "",
			Message:           "Amostra criada com sucesso",
			HttpStatusMessage: "OK",
			HttpStatusCode:    200,
		},
	})
}

func UpdateSample(c *gin.Context) {
	var sample entities.Sample
	if err := c.ShouldBindJSON(&sample); err != nil {
		c.JSON(500, gin.H{
			"retorno": models.ReturnModel{
				Trace:             "",
				Message:           "Erro ao atualizar amostra: " + err.Error(),
				HttpStatusMessage: "Internal Server Error",
				HttpStatusCode:    500,
			},
		})
		return
	}
	repositories.UpdateSample(sample.ID, sample)
	c.JSON(http.StatusOK, gin.H{
		"updatedSample": sample,
		"retorno": models.ReturnModel{
			Trace:             "",
			Message:           "Amostra atualizada com sucesso",
			HttpStatusMessage: "OK",
			HttpStatusCode:    200,
		},
	})
}

func DeleteSample(c *gin.Context) {
	var idFromInput uint64
	idFromInput, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"retorno": models.ReturnModel{
				Trace:             "",
				Message:           "Erro ao deletar amostra: " + err.Error(),
				HttpStatusMessage: "Internal Server Error",
				HttpStatusCode:    500,
			},
		})
		return
	}
	repositories.DeleteSampleById(idFromInput)
	c.JSON(http.StatusOK, gin.H{
		"retorno": models.ReturnModel{
			Trace:             "",
			Message:           "Amostra deletada com sucesso",
			HttpStatusMessage: "OK",
			HttpStatusCode:    200,
		},
	})
}
