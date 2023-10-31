package models

import (
	_ "gorm.io/driver/postgres"
	"time"
)

type Modificacao struct {
	DataEmenta *time.Time `gorm:"primaryKey" json:"dataEmenta"`
	DataGeral  *time.Time `gorm:"primaryKey" json:"dataGeral"`
}

func (Modificacao) TableName() string {
	return "modificacao"
}
