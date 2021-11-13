package main

import (
	"github.com/enestkn/golangstudy1/labs/lab3/config"
	"github.com/enestkn/golangstudy1/labs/lab3/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	r := gin.Default()
	r.GET("/ping", controller.GETPing)

	r.POST("/personel", controller.POSTPersonal)
	r.GET("/personel/1", controller.GETPersonalByID)
	r.GET("personels", controller.GETPersonals)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
