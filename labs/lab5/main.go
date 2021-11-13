package main

import (
	"github.com/enestkn/golangstudy1/labs/lab5/config"
	"github.com/enestkn/golangstudy1/labs/lab5/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	r := gin.Default()
	r.GET("/ping", controller.GETPing)

	// Personal Rooter
	r.POST("/personel", controller.POSTPersonal)
	r.GET("personels", controller.GETPersonals)
	r.GET("/personel/:id", controller.GETPersonalByID)
	r.DELETE("/personel/:id", controller.DELETEPersonal)
	r.PUT("/personel/:id", controller.UPDATEPersonalByID)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
