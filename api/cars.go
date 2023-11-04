package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mcabej/db"
	"github.com/mcabej/db/models"
	"github.com/mcabej/helpers"
	"github.com/mcabej/services"
)

type CarRequest struct {
	Make      string    `json:"make" binding:"required"`
	Model     string    `json:"model" binding:"required"`
	BuildDate time.Time `json:"buildDate" binding:"required"`
	ColorID   uint      `json:"color" binding:"required"`
}

type CarResponse struct {
	ID        uint      `json:"id"`
	Make      string    `json:"make" binding:"required"`
	Model     string    `json:"model" binding:"required"`
	BuildDate time.Time `json:"buildDate" binding:"required"`
	Color     string    `json:"color"`
	ColorID   uint      `json:"colorId" binding:"required"`
}

func CreateCar(c *gin.Context) {
	var req CarRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	car := models.Car{
		Make:      req.Make,
		Model:     req.Model,
		BuildDate: req.BuildDate,
		ColorID:   req.ColorID,
	}

	// create car
	err := services.CreateCar(&car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// check that color exist
	color, err := helpers.ValidateColorExist(int(req.ColorID))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	carWithColorName := CarResponse{
		ID:        car.ID,
		Make:      car.Make,
		Model:     car.Model,
		BuildDate: car.BuildDate,
		Color:     color.Name,
		ColorID:   color.ID,
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": carWithColorName,
	})
}

func ListCars(c *gin.Context) {
	var cars []models.Car

	results := db.DB.Find(&cars)
	if results.Error != nil {
		c.JSON(http.StatusInternalServerError, results.Error.Error())
		return
	}

	var carsWithColorNames []CarResponse
	for _, value := range cars {
		var color models.Color
		result := db.DB.First(&color, value.ColorID)
		if result.Error != nil {
			continue
		}

		carsWithColorNames = append(
			carsWithColorNames,
			CarResponse{
				ID:        value.ID,
				Make:      value.Make,
				Model:     value.Model,
				BuildDate: value.BuildDate,
				Color:     color.Name,
				ColorID:   color.ID,
			},
		)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": carsWithColorNames,
	})
}

func GetCar(c *gin.Context) {
	carId := c.Param("id")

	car := &models.Car{}
	err := services.GetCar(carId, car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": car,
	})
}

type UpdateCarRequest struct {
	Make      string    `json:"make"`
	Model     string    `json:"model"`
	BuildDate time.Time `json:"buildDate"`
	ColorID   uint      `json:"color"`
}

func UpdateCar(c *gin.Context) {
	// find car
	carId := c.Param("id")

	// get req
	var req UpdateCarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// update car
	var car = models.Car{
		Make:      req.Make,
		Model:     req.Model,
		BuildDate: req.BuildDate,
		ColorID:   req.ColorID,
	}

	if err := services.UpdateCar(carId, &car); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	color, err := helpers.ValidateColorExist(int(req.ColorID))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	carWithColorName := CarResponse{
		ID:        car.ID,
		Make:      car.Make,
		Model:     car.Model,
		BuildDate: car.BuildDate,
		Color:     color.Name,
		ColorID:   color.ID,
	}

	c.JSON(http.StatusOK, gin.H{
		"result": carWithColorName,
	})
}

func DeleteCar(c *gin.Context) {
	carId := c.Param("id")

	err := services.DeleteCar(carId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
