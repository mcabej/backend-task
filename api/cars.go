package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mcabej/db"
	"github.com/mcabej/db/models"
	"github.com/mcabej/helpers"
)

type carRequest struct {
	Make      string    `json:"make" binding:"required"`
	Model     string    `json:"model" binding:"required"`
	BuildDate time.Time `json:"build_date" binding:"required"`
	ColorID   uint      `json:"color" binding:"required"`
}

func CreateCar(c *gin.Context) {
	var req carRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check that age of car is not older than four years
	if err := helpers.ValidateCarAge(req.BuildDate); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check that color exist
	if err := helpers.ValidateColorExist(int(req.ColorID)); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	car := models.Car{
		Make:      req.Make,
		Model:     req.Model,
		BuildDate: req.BuildDate,
		ColorID:   req.ColorID,
	}

	result := db.DB.Create(&car)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": car,
	})
}

func ListCars(c *gin.Context) {
	var cars []models.Car

	results := db.DB.Find(&cars)

	if results.Error != nil {
		c.JSON(http.StatusInternalServerError, results.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": cars,
	})
}

func GetCar(c *gin.Context) {
	carId := c.Param("id")

	var car models.Car
	result := db.DB.First(&car, carId)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": car,
	})
}

type updateCarRequest struct {
	Make      string     `json:"make"`
	Model     string     `json:"model"`
	BuildDate *time.Time `json:"build_date,omitempty"`
	ColorID   uint       `json:"color"`
}

func UpdateCar(c *gin.Context) {
	// find car
	carId := c.Param("id")

	var car models.Car

	result := db.DB.First(&car, carId)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	// get req
	var req updateCarRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check that age of car is not older than four years
	if err := helpers.ValidateCarAge(*req.BuildDate); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check that color exist
	if err := helpers.ValidateColorExist(int(req.ColorID)); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// update car
	db.DB.Model(&car).Updates(models.Car{
		Make:      req.Make,
		Model:     req.Model,
		BuildDate: *req.BuildDate,
		ColorID:   req.ColorID,
	})

	c.JSON(http.StatusOK, gin.H{
		"result": car,
	})
}

func DeleteCar(c *gin.Context) {
	carId := c.Param("id")

	db.DB.Delete(&models.Car{}, carId)

	c.Status(http.StatusOK)
}
