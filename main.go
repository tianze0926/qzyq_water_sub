package main

import (
	"embed"
	"flag"
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/tz039e/water_sub/controllers"
	"github.com/tz039e/water_sub/models"
)

//go:embed frontend/build
var frontendFS embed.FS

func main() {
	// flags
	addr := flag.String("addr", "", "listen address")
	port := flag.String("port", "", "port")
	flag.Parse()
	if *port == "" {
		panic("port is empty")
	}

	// gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// routing
	r.Use(static.Serve("/", EmbedDir(frontendFS, "frontend/build")))
	api := r.Group("/api")
	{
		api.GET("/record", controllers.GetRecords)
		api.POST("/record", controllers.PostRecord)
	}

	// db setup
	models.ConnectDB()

	// run
	fmt.Println("listening on port", *port)
	r.Run(*addr + ":" + *port)
}
