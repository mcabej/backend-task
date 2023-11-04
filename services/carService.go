package services

import (
	"fmt"

	"github.com/mcabej/db"
	"github.com/mcabej/db/models"
	"github.com/mcabej/helpers"
)

func CreateCar(car *models.Car) error {
	// check that age of car is not older than four years
	if err := helpers.ValidateCarAge(car.BuildDate); err != nil {
		return err
	}

	// check that color exist
	_, err := helpers.ValidateColorExist(int(car.ColorID))
	if err != nil {
		return err
	}

	// create car
	result := db.DB.Create(&car)
	if result.Error != nil {
		return err
	}

	return nil
}

func UpdateCar(id string, car *models.Car) error {
	// check that age of car is not older than four years
	if err := helpers.ValidateCarAge(car.BuildDate); err != nil {
		return err
	}

	// check that color exist
	_, err := helpers.ValidateColorExist(int(car.ColorID))
	if err != nil {
		return err
	}

	// start a transaction
	tx := db.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// find car to update
	var carToUpdate models.Car

	result := tx.First(&carToUpdate, id)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// update car
	result = tx.Model(&carToUpdate).Updates(models.Car{
		Make:      car.Make,
		Model:     car.Model,
		BuildDate: car.BuildDate,
		ColorID:   car.ColorID,
	})

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	// Return the updated car
	*car = carToUpdate

	return nil
}

func DeleteCar(id string) error {
	result := db.DB.Delete(&models.Car{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetCar(id string, car *models.Car) error {
	result := db.DB.First(&car, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
