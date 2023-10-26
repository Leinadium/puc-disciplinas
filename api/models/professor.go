package models

import (
	_ "gorm.io/driver/postgres"
)

type Professor struct {
	NomeProfessor string `gorm:"primaryKey"`
}
