package controller

import (
	"github.com/enestkn/golangstudy1/labs/lab3/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func POSTPersonal(c *gin.Context) {
	var form config.Form
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	config.Personels = append(config.Personels, form)
	c.JSON(http.StatusCreated, form)
}

func GETPersonalByID(c *gin.Context) {
	c.JSON(http.StatusOK, config.Personels[0])
}

func GETPersonals(c *gin.Context) {
	c.JSON(http.StatusOK, config.Personels)
}
