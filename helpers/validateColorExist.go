package helpers

import (
	"errors"

	"github.com/mcabej/db"
	"github.com/mcabej/db/models"
)

func ValidateColorExist(colorId int) (*models.Color, error) {
	var color models.Color

	result := db.DB.First(&color, colorId)

	if result.Error != nil || result.RowsAffected <= 0 {
		return nil, errors.New("there's no such color")
	}

	return &color, nil
}
