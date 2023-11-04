package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mcabej/db"
	"github.com/mcabej/db/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectToDB()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func createMockCar(t *testing.T) models.Car {
	newCar := models.Car{
		Make:      "Tesla",
		Model:     "Model X",
		BuildDate: time.Now(),
		ColorID:   1,
	}

	tx := db.DB.Create(&newCar)
	require.NoError(t, tx.Error)

	return newCar
}

func TestCreateCar(t *testing.T) {
	router := setupRouter()
	router.POST("api/car/create")

	newCar := models.Car{
		Make:      "Honda",
		Model:     "Prius",
		BuildDate: time.Now(),
		ColorID:   1,
	}

	jsonValue, _ := json.Marshal(newCar)
	req, _ := http.NewRequest("POST", "api/car/create", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetCar(t *testing.T) {
	router := setupRouter()
	router.GET("api/car/:id")

	mockCar := createMockCar(t)

	w := httptest.NewRecorder()
	url := fmt.Sprintf("/api/car/%d", mockCar.ID)
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
