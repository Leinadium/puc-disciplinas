package models

type Semestre struct {
	CodDisciplina string `gorm:"primaryKey"`
	CodCurriculo  uint8  `gorm:"primaryKey"`
	Semestre      uint8
}
