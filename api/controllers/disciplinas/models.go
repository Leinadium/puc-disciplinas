package disciplinas

import (
	"github.com/Leinadium/puc-disciplinas/api/controllers"
)

type ResultCodigo struct {
	CodDisciplina string `json:"codigo" gorm:"column:cod_disciplina"`
}

type ResultLista struct {
	CodDisciplina  string `json:"codigo" gorm:"column:cod_disciplina"`
	NomeDisciplina string `json:"nome" gorm:"column:nome_disciplina"`
	Creditos       int    `json:"creditos" gorm:"column:creditos"`
	QtdVagas       int    `json:"qtdVagas" gorm:"column:qtdvagas"`
	QtdTurmas      int    `json:"qtdTurmas" gorm:"column:qtdturmas"`
}

type ResultInfo struct {
	CodDisciplina  string  `json:"codigo" gorm:"column:cod_disciplina"`
	NomeDisciplina string  `json:"nome" gorm:"column:nome_disciplina"`
	Creditos       int     `json:"creditos" gorm:"column:creditos"`
	Ementa         string  `json:"ementa" gorm:"column:ementa"`
	GrupoPrereq    int     `json:"grupoPrereq" gorm:"column:grupo_prereq"`
	CodDiscDepen   string  `json:"codDiscDepen" gorm:"column:cod_disc_depen"`
	MediaGrau      float64 `json:"mediaGrau" gorm:"column:media_grau"`
	QtdAlunos      int     `json:"qtdAlunos" gorm:"column:qtd_alunos"`
	MediaAvaliacao float64 `json:"mediaAvaliacao" gorm:"column:media_avaliacao"`
	QtdAvaliacoes  int     `json:"qtdAvaliacoes" gorm:"column:qtd_avaliacoes"`
}

type ResultTurmas struct {
	CodTurma      string  `json:"codigo" gorm:"column:cod_turma"`
	NomeProfessor string  `json:"nomeProfessor" gorm:"column:nome_professor"`
	NotaProfessor float64 `json:"notaProfessor" gorm:"column:nota_professor"`
	Shf           int     `json:"shf" gorm:"column:shf"`
	Dia           string  `json:"dia" gorm:"column:dia"`
	HoraInicio    int     `json:"horaInicio" gorm:"column:hora_inicio"`
	HoraFim       int     `json:"horaFim" gorm:"column:hora_fim"`
	Destino       string  `json:"destino" gorm:"column:destino"`
	Vagas         int     `json:"vagas" gorm:"column:vagas"`
}

type RespostaDisciplina struct {
	CodDisciplina  string                   `json:"codigo"`
	NomeDisciplina string                   `json:"nome"`
	Ementa         string                   `json:"ementa"`
	Creditos       int                      `json:"creditos"`
	PreRequisitos  []RespostaPreReqs        `json:"preRequisitos"`
	GrauMedio      controllers.RoundedFloat `json:"grauMedio"`
	QtdAlunos      int                      `json:"qtdAlunos"`
	AvaliacaoMedia controllers.RoundedFloat `json:"avaliacaoMedia"`
	QtdAvaliacoes  int                      `json:"qtdAvaliacoes"`
	Turmas         []RespostaTurmas         `json:"turmas"`
}

type RespostaPreReqs struct {
	GrupoId int      `json:"grupoId"`
	PreReqs []string `json:"preReqs"`
}

type RespostaTurmas struct {
	CodTurma      string                   `json:"codigo"`
	NomeProfessor string                   `json:"professor"`
	NotaProfessor controllers.RoundedFloat `json:"notaProfessor"`
	Shf           int                      `json:"shf"`
	Horarios      []RespostaHorarios       `json:"horarios"`
	Alocacoes     []RespostaAlocacoes      `json:"alocacoes"`
}

type RespostaHorarios struct {
	Dia        string `json:"dia"`
	HoraInicio int    `json:"inicio"`
	HoraFim    int    `json:"fim"`
}

type RespostaAlocacoes struct {
	Destino  string `json:"destino"`
	QtdVagas int    `json:"vagas"`
}
