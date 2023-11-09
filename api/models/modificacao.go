package models

import (
	_ "gorm.io/driver/postgres"
	"time"
)

type Modificacao struct {
	DataEmenta   *time.Time `gorm:"primaryKey" json:"dataEmenta"`
	DataGeral    *time.Time `gorm:"primaryKey" json:"dataGeral"`
	ModoFallback bool       `gorm:"primaryKey" json:"modoFallback"`
}

func (Modificacao) TableName() string {
	return "modificacao"
}
