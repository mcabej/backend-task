package initialise

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("APP_ENV") != "prod" {
		godotenv.Load() // The Original .env
	}
}
