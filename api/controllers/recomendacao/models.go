package recomendacao

import "github.com/Leinadium/puc-disciplinas/api/controllers"

type InputRecomendacao struct {
	Escolhas []InputEscolha `json:"escolhas"`
}

type InputEscolha struct {
	CodDisciplina string `json:"disciplina"`
	CodTurma      string `json:"turma"`
}

type resultQuery struct {
	CodDisciplina string  `gorm:"column:cod_disciplina"`
	Conteudo1     bool    `gorm:"column:conteudo1"`
	Conteudo21    int     `gorm:"column:conteudo21"`
	Conteudo22    int     `gorm:"column:conteudo22"`
	Horario1      int     `gorm:"column:horario1"`
	Horario2      int     `gorm:"column:horario2"`
	Opiniao       float64 `gorm:"column:opiniao"`
	Professor     float64 `gorm:"column:professor"`
	Avaliacao     float64 `gorm:"column:avaliacao"`
	Existe        bool    `gorm:"column:existe"`
}

type ResultAlg struct {
	CodDisciplina string                   `json:"cod"`
	Valor         controllers.RoundedFloat `json:"val"`
	Pesos         ResultPesos              `json:"pes"`
}

type ResultPesos struct {
	Conteudo  controllers.RoundedFloat `json:"c"`
	Horario   controllers.RoundedFloat `json:"h"`
	Opiniao   controllers.RoundedFloat `json:"o"`
	Professor controllers.RoundedFloat `json:"p"`
	Avaliacao controllers.RoundedFloat `json:"a"`
}
