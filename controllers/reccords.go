package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
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

const API_URL = "http://dingshui.bjqzhd.com"

func PostRecord(c *gin.Context) {
	var input models.Record
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	input.IP = c.ClientIP()

	// start transaction
	tx := models.DB.Begin()
	// add to db
	tx.Create(&input)
	// reqeust to api
	// step 1: get user info
	resp, err := http.PostForm(API_URL+"/auser/getuser.html", url.Values{
		"param": {input.FileNumber},
		"name":  {"pw"},
	})
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var userInfo map[string]string
	json.NewDecoder(resp.Body).Decode(&userInfo)
	// step 2: submit
	resp, err = http.PostForm(API_URL+"/buy/subs.html", url.Values{
		"pw":      {userInfo["pw"]},
		"name":    {userInfo["name"]},
		"num":     {"1"},
		"num1":    {"0"},
		"lid":     {"6"},
		"phone":   {""},
		"address": {""},
	})
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	if resp.StatusCode != 200 {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"msg": "sub returned code: " + strconv.Itoa(resp.StatusCode)})
		return
	}
	// success
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"data": input,
	})
}
