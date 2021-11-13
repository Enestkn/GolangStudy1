package controller

import (
	"errors"
	"github.com/enestkn/golangstudy1/labs/lab4/models"
	"github.com/enestkn/golangstudy1/labs/lab5/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// POSTPersonal Insert
func POSTPersonal(c *gin.Context) {
	var form models.Personal

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := config.DB.Create(&form).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, form)
}

// GETPersonalByID Get Personal By ID
func GETPersonalByID(c *gin.Context) {
	id := c.Param("id")
	var personal models.Personal
	if err := config.DB.First(&personal, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, personal)
}

func GETPersonals(c *gin.Context) {
	var personals []models.Personal
	if err := config.DB.Find(&personals).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, personals)
}

func DELETEPersonal(c *gin.Context) {
	id := c.Param("id")
	var personal models.Personal
	if err := config.DB.First(&personal, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, err)
		return
	}
	config.DB.Delete(&personal)
	c.JSON(http.StatusOK, "deleted")
}

func UPDATEPersonalByID(c *gin.Context) {
	id := c.Param("id")
	var form models.Personal
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var personal models.Personal
	if err := config.DB.First(&personal, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, err)
		return
	}

	config.DB.Save(&form)

	c.JSON(http.StatusOK, form)
}
