package avaliacoes

import "gopkg.in/guregu/null.v3"

type ItemProfessorCompleto struct {
	Nome       string     `json:"nome" gorm:"column:nome_professor"`
	Nota       null.Int   `json:"nota" gorm:"column:nota_avaliacao"`
	Media      null.Float `json:"media" gorm:"column:media"`
	Quantidade int32      `json:"qtd" gorm:"column:quantidade"`
}

type ItemDisciplinaCompleta struct {
	Nome       string     `json:"nome" gorm:"column:nome_disciplina"`
	Codigo     string     `json:"codigo" gorm:"column:cod_disciplina"`
	Nota       null.Int   `json:"nota" gorm:"column:nota_avaliacao"`
	Media      null.Float `json:"media" gorm:"column:media"`
	Quantidade int32      `json:"qtd" gorm:"column:quantidade"`
}

type ResultGetAvaliacao struct {
	Professores []ItemProfessorCompleto  `json:"professores"`
	Disciplinas []ItemDisciplinaCompleta `json:"disciplinas"`
}
