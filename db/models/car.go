package models

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Make      string
	Model     string
	BuildDate time.Time
	ColorID   uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
