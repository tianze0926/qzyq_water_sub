package main

import (
	"github.com/gin-gonic/gin"

	"github.com/tz039e/water_sub/controllers"
	"github.com/tz039e/water_sub/models"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/record", controllers.GetRecords)
		api.POST("/record", controllers.PostRecord)
	}

	models.ConnectDB()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
