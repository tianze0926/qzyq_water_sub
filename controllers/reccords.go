package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tz039e/water_sub/models"
)

func GetRecords(c *gin.Context) {
	var records []models.Record
	models.DB.Order("id desc").Limit(10).Find(&records)

	c.JSON(http.StatusOK, gin.H{
		"data": records,
	})
}

func PostRecord(c *gin.Context) {
	var input models.Record
	if err := c.BindJSON(&input); err != nil {
		return
	}

	models.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{
		"data": input,
	})
}
