package controllers

import "github.com/gin-gonic/gin"

func GetInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"sample": "Hello World",
	})
}
