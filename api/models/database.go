package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func GetDB() (*gorm.DB, error) {
	var err error = nil
	if db == nil {
		// cria se nao existir
		dsn := os.Getenv("GORM_DSN")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	return db, err
}
