package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tz039e/water_sub/models"
)

func GetRecords(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "param `page` error"})
		return
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "param `size` error"})
		return
	}

	var records []models.Record
	models.DB.Order("id desc").Limit(size).Offset(size * page).Find(&records)

	var count int64
	models.DB.Model(&models.Record{}).Count(&count)

	c.JSON(http.StatusOK, gin.H{
		"data":  records,
		"total": count,
	})
}

func PostRecord(c *gin.Context) {
	var input models.Record
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	input.IP = c.ClientIP()

	models.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{
		"data": input,
	})
}
