package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabaseConnection() (*gorm.DB, error) {
	pgConfig := postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=password dbname=books_db port=12345 sslmode=disable TimeZone=Europe/Bucharest",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	})
	db, err := gorm.Open(pgConfig, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
