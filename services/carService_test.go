package services

import (
	"strconv"
	"testing"
	"time"

	"github.com/mcabej/db"
	"github.com/mcabej/db/models"
	"github.com/mcabej/initialise"
	"github.com/stretchr/testify/assert"
)

func init() {
	initialise.LoadEnv()
	db.ConnectToDB()
}

func TestCreateValidCar(t *testing.T) {
	newCar := models.Car{
		Make:      "Tesla",
		Model:     "Model X",
		BuildDate: time.Now(),
		ColorID:   1,
	}

	err := CreateCar(&newCar)

	assert.NoError(t, err)
}

func TestCreateInvalidCarBuildDate(t *testing.T) {
	newCar := models.Car{
		Make:      "Tesla Failed Build Random",
		Model:     "Model XYZ",
		BuildDate: time.Now().AddDate(4, 0, 0),
		ColorID:   1,
	}

	err := CreateCar(&newCar)
	assert.Error(t, err)

	var queryResults []models.Car
	db.DB.Where("make = ?", newCar.Make).Find(&queryResults)

	assert.Empty(t, queryResults)
}

func TestUpdateCar(t *testing.T) {
	car := models.Car{
		Make:      "Tesla",
		Model:     "Model",
		BuildDate: time.Now(),
		ColorID:   1,
	}

	err := CreateCar(&car)
	assert.NoError(t, err)

	newCar := models.Car{
		Make:      "Ferrari",
		Model:     "Model",
		BuildDate: time.Now(),
		ColorID:   1,
	}

	stringId := strconv.FormatUint(uint64(car.ID), 10)

	err = UpdateCar(stringId, &newCar)
	assert.NoError(t, err)

	assert.Equal(t, car.ID, newCar.ID)
	assert.NotEqual(t, car.Make, newCar.Make)
}
