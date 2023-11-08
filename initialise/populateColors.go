package initialise

import (
	"log"

	"github.com/mcabej/db"
	"github.com/mcabej/db/models"
)

func PopulateColors() {
	colors := []*models.Color{
		{Name: "red"},
		{Name: "blue"},
		{Name: "white"},
		{Name: "black"},
	}

	result := db.DB.Create(colors)
	if result.Error != nil {
		log.Fatal("Unable to create default colors")
	}
}
