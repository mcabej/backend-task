package main

import (
	"github.com/gin-contrib/static"
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

	router.Use(static.Serve("/", static.LocalFile("./app/build", true)))

	car := router.Group("api/car")
	{
		car.GET("/:id", api.GetCar)
		car.GET("/cars", api.ListCars)

		car.POST("/create", api.CreateCar)
		car.PUT("/:id", api.UpdateCar)

		car.DELETE("/:id", api.DeleteCar)
	}

	color := router.Group("api/color")
	{
		color.GET("/:id", api.GetColor)
		color.GET("/colors", api.ListColors)

		color.POST("/create", api.CreateColor)
		color.PUT("/:id", api.UpdateColor)

		color.DELETE("/:id", api.DeleteColor)
	}

	router.Run()
}
