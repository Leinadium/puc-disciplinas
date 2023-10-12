package models

import (
	_ "gorm.io/driver/postgres"
)

type Historico struct {
	CodUsuario    string `gorm:"primaryKey"`
	CodDisciplina string `gorm:"primaryKey"`
	Semestre      int    `gorm:"primaryKey"`
	Grau          int
}
