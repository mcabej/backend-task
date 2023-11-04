package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcabej/db"
	"github.com/mcabej/db/models"
)

type colorRequest struct {
	Name string `json:"make" binding:"required"`
}

func CreateColor(c *gin.Context) {
	var req colorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	color := models.Color{
		Name: req.Name,
	}

	result := db.DB.Create(&color)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": color,
	})
}

func ListColors(c *gin.Context) {
	var colors []models.Color

	results := db.DB.Find(&colors)

	if results.Error != nil {
		c.JSON(http.StatusInternalServerError, results.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": colors,
	})
}

func GetColor(c *gin.Context) {
	id := c.Param("id")

	var color models.Color

	result := db.DB.First(&color, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": color,
	})
}

func UpdateColor(c *gin.Context) {
	id := c.Param("id")

	var color models.Color

	result := db.DB.First(&color, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	// get req
	var req colorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	db.DB.Model(&color).Updates(models.Color{
		Name: req.Name,
	})

	c.JSON(http.StatusOK, gin.H{
		"result": color,
	})
}

func DeleteColor(c *gin.Context) {
	id := c.Param("id")

	db.DB.Delete(&models.Color{}, id)

	c.Status(http.StatusOK)
}
