package models

import (
	"database/sql"
	_ "gorm.io/driver/postgres"
)

type Usuario struct {
	CodUsuario   string `gorm:"primaryKey"`
	NomeUsuario  string
	CodCurriculo sql.NullString
}
