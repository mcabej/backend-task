package helpers

import (
	"errors"

	"github.com/mcabej/db"
	"github.com/mcabej/db/models"
)

func ValidateColorExist(colorId int) error {
	var color models.Color

	result := db.DB.First(&color, colorId)

	if result.Error != nil || result.RowsAffected <= 0 {
		return errors.New("there's no such color")
	}

	return nil
}
