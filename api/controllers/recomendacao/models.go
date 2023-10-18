package recomendacao

type inputRecomendacao struct {
	Escolhas []inputEscolha `json:"escolhas"`
}

type inputEscolha struct {
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
}

type resultAlg struct {
	CodDisciplina string  `json:"cod"`
	Nota          float64 `json:"nota"`
}
