package db

import (
	"time"
)

type Book struct {
	ID          int       `gorm:"primaryKey"`
	Title       string    `gorm:"size:255;not null"`
	AuthorID    int       `gorm:"not null"`
	PublishDate time.Time `gorm:"not null"`
}
