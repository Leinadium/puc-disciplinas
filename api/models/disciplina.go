package models

import (
	"database/sql"
	_ "gorm.io/driver/postgres"
)

type Disciplina struct {
	CodDisciplina  string `gorm:"primaryKey"`
	CodDepto       string
	NomeDisciplina string
	Ementa         sql.NullString
	Creditos       uint8
}
