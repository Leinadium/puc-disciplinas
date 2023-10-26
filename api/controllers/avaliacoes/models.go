package avaliacoes

import "gopkg.in/guregu/null.v3"

type itemProfessorCompleto struct {
	Nome       string     `json:"nome" gorm:"column:nome_professor"`
	Nota       null.Int   `json:"nota" gorm:"column:nota_avaliacao"`
	Media      null.Float `json:"media" gorm:"column:media"`
	Quantidade int32      `json:"qtd" gorm:"column:quantidade"`
}

type itemDisciplinaCompleta struct {
	Nome       string     `json:"nome" gorm:"column:nome_disciplina"`
	Codigo     string     `json:"codigo" gorm:"column:cod_disciplina"`
	Nota       null.Int   `json:"nota" gorm:"column:nota_avaliacao"`
	Media      null.Float `json:"media" gorm:"column:media"`
	Quantidade int32      `json:"qtd" gorm:"column:quantidade"`
}

type resultGetAvaliacao struct {
	Professores []itemProfessorCompleto  `json:"professores"`
	Disciplinas []itemDisciplinaCompleta `json:"disciplinas"`
}
