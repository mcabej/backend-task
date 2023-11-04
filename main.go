package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mcabej/api"
	"github.com/mcabej/db"
	"github.com/mcabej/initialise"
)

func init() {
	initialise.LoadEnv()
	db.ConnectToDB()
}

func main() {
	router := gin.Default()

	car := router.Group("/car")
	{
		car.GET("/:id", api.GetCar)
		car.GET("/list", api.ListCars)

		car.POST("/create", api.CreateCar)
		car.PUT("/:id", api.UpdateCar)

		car.DELETE("/:id", api.DeleteCar)
	}

	color := router.Group("/color")
	{
		color.GET("/:id", api.GetColor)
		color.GET("/list", api.ListColors)

		color.POST("/create", api.CreateColor)
		color.PUT("/update/:id", api.UpdateColor)

		color.DELETE("/delete/:id", api.DeleteColor)
	}

	router.Run()
}
