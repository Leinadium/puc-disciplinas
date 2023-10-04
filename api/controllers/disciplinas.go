package controllers

import (
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDisciplinas(c *gin.Context) {
	// pega o db
	var db = GetDbOrSetError(c)
	if db == nil {
		return
	}

	var disciplinas []models.Disciplina
	db.Find(&disciplinas)

	c.JSON(http.StatusOK, gin.H{"data": disciplinas})
}

func GetDisciplinasResumo(c *gin.Context) {
	// pega o db
	var db = GetDbOrSetError(c)
	if db == nil {
		return
	}

	type result struct {
		CodDisciplina  string
		NomeDisciplina string
		Creditos       uint8
		QtdVagas       uint8
		QtdTurmas      uint8
	}

	// var results []result

	//SELECT d.cod_disciplina, d.nome_disciplina, d.creditos, COUNT(t.cod_turma) qtdTurmas, SUM(t.vagas) qtdVagas
	//FROM disciplinas d
	//LEFT JOIN (
	//	SELECT t.cod_turma, t.cod_disciplina, SUM(a.vagas) vagas
	//FROM turmas t LEFT JOIN alocacoes a on t.cod_turma = a.cod_turma and t.cod_disciplina = a.cod_disciplina
	//GROUP BY t.cod_turma, t.cod_disciplina
	//) t
	//ON d.cod_disciplina = t.cod_disciplina
	//GROUP BY d.cod_disciplina, d.nome_disciplina, d.creditos

}
