package models

type Curriculo struct {
	CodCurriculo  uint8 `gorm:"primaryKey"`
	NomeCurriculo string
}
