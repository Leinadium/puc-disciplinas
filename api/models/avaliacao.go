package models

import (
	_ "gorm.io/driver/postgres"
	"time"
)

type AvaliacaoDisciplina struct {
	CodUsuario    string `gorm:"primaryKey"`
	CodDisciplina string `gorm:"primaryKey"`
	NotaAvaliacao uint8
	DataAvaliacao *time.Time
}

type AvaliacaoProfessor struct {
	CodUsuario    string `gorm:"primaryKey"`
	NomeProfessor string `gorm:"primaryKey"`
	NotaAvaliacao uint8
	DataAvaliacao *time.Time
}

func (AvaliacaoDisciplina) TableName() string {
	return "avaliacoes_disciplinas"
}

func (AvaliacaoProfessor) TableName() string {
	return "avaliacoes_professores"
}
