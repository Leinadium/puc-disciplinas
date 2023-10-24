package models

import (
	_ "gorm.io/driver/postgres"
)

type Grade struct {
	CodGrade   string `gorm:"primaryKey"`
	CodUsuario string
	Conteudo   string
}
