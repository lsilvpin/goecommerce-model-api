package controllers

import "github.com/gin-gonic/gin"

func GetSamples(c *gin.Context) {
	c.JSON(200, gin.H{
		"sample": "Hello World",
	})
}

func GetSampleById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"sampleId": id,
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
