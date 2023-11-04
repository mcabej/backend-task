package helpers

import (
	"errors"
	"time"

	age "github.com/theTardigrade/golang-age"
)

func ValidateCarAge(buildDate time.Time) error {
	carAge := age.CalculateToNow(buildDate)

	if carAge > 4 {
		return errors.New("car cannot be older than four years")
	}

	if carAge < 0 {
		return errors.New("build date cannot be in the future")
	}

	return nil
}
