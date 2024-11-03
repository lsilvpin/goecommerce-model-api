package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	entities "github.com/lsilvpin/goecommerce-model-api/main/library/domain/entities"
	models "github.com/lsilvpin/goecommerce-model-api/main/library/domain/models"
	"github.com/lsilvpin/goecommerce-model-api/main/library/repositories"
)

func GetSamples(c *gin.Context) {
	samples := repositories.ReadAllSamples()
	length := len(samples)
	c.JSON(200, gin.H{
		"length":  length,
		"samples": samples,
		"retorno": models.ReturnModel{
			Trace:   "",
			Message: "Amostras retornadas com sucesso",
		},
	})
}

func GetSampleById(c *gin.Context) {
	var idFromInput uint64
	idFromInput, err := strconv.ParseUint(c.Param("id"), 10, 64)
	log.Println("idFromInput: ", idFromInput)
	if err != nil {
		c.JSON(500, gin.H{
			"retorno": models.ReturnModel{
				Trace:   "",
				Message: "Erro ao buscar amostra: " + err.Error(),
			},
		})
		return
	}
	sample, err := repositories.ReadSampleById(idFromInput)
	if err != nil {
		c.JSON(404, gin.H{
			"retorno": models.ReturnModel{
				Trace:   "",
				Message: "Amostra de id " + strconv.FormatUint(idFromInput, 10) + " não encontrada",
			},
		})
		return
	}
	c.JSON(200, gin.H{
		"sample": sample,
		"retorno": models.ReturnModel{
			Trace:   "",
			Message: "Amostra retornada com sucesso",
		},
	})
}

func CreateSample(c *gin.Context) {
	var sample entities.Sample
	if err := c.ShouldBindJSON(&sample); err != nil {
		c.JSON(500, gin.H{
			"retorno": models.ReturnModel{
				Trace:   "",
				Message: "Erro ao criar amostra: " + err.Error(),
			},
		})
		return
	}
	createdId := repositories.CreateSample(sample)
	createdSample, createErr := repositories.ReadSampleById(createdId)
	if createErr != nil {
		c.JSON(500, gin.H{
			"retorno": models.ReturnModel{
				Trace:   "",
				Message: "Erro após criar sample: " + createErr.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"createdSample": createdSample,
		"retorno": models.ReturnModel{
			Trace:   "",
			Message: "Amostra criada com sucesso",
		},
	})
}

func UpdateSample(c *gin.Context) {
	idFromInput, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"retorno": models.ReturnModel{
				Trace:   "",
				Message: "Erro ao obter id da amostra: " + err.Error(),
			},
		})
		return
	}
	var sample entities.Sample
	if err := c.ShouldBindJSON(&sample); err != nil {
		c.JSON(500, gin.H{
			"retorno": models.ReturnModel{
				Trace:   "",
				Message: "Erro ao obter amostra para atualização: " + err.Error(),
			},
		})
		return
	}
	sampleAfterUpdate, updateErr := repositories.UpdateSample(idFromInput, sample)
	if updateErr != nil {
		c.JSON(404, gin.H{
			"retorno": models.ReturnModel{
				Trace:   "",
				Message: "Amostra de id " + strconv.FormatUint(idFromInput, 10) + " não encontrada",
			},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"updatedSample": sampleAfterUpdate,
		"retorno": models.ReturnModel{
			Trace:   "",
			Message: "Amostra atualizada com sucesso",
		},
	})
}

func DeleteSample(c *gin.Context) {
	var idFromInput uint64
	idFromInput, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"retorno": models.ReturnModel{
				Trace:   "",
				Message: "Erro ao deletar amostra: " + err.Error(),
			},
		})
		return
	}
	deleteErr := repositories.DeleteSampleById(idFromInput)
	if deleteErr != nil {
		c.JSON(404, gin.H{
			"retorno": models.ReturnModel{
				Trace:   "",
				Message: "Amostra de id " + strconv.FormatUint(idFromInput, 10) + " não encontrada",
			},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"retorno": models.ReturnModel{
			Trace:   "",
			Message: "Amostra de id " + strconv.FormatUint(idFromInput, 10) + " deletada com sucesso",
		},
	})
}
