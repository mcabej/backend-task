package main

import (
	"github.com/mcabej/db"
	"github.com/mcabej/db/models"
	"github.com/mcabej/initialise"
)

func init() {
	initialise.LoadEnv()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(&models.Car{}, &models.Color{})
	initialise.PopulateColors()
}
