package db

import (
	"time"
)

type Author struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null"`
	DateOfBirth time.Time `gorm:"not null"`
}
