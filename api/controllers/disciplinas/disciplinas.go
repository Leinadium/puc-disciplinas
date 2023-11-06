package disciplinas

import (
	"database/sql"
	"github.com/Leinadium/puc-disciplinas/api/controllers"
	"github.com/Leinadium/puc-disciplinas/api/controllers/queries"
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type resultCompleto struct {
	Disciplinas []resultLista      `json:"disciplinas"`
	Modificacao models.Modificacao `json:"modificacao"`
}

func GetDisciplinasInformacoes(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	// faz a query
	var results []resultLista
	err := db.Raw(queries.QUERY_INFORMACOES).Scan(&results).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}

	// pega as modificacoes (se der erro, tudo bem)
	var modificacao models.Modificacao
	db.First(&modificacao)

	completo := resultCompleto{
		Disciplinas: results,
		Modificacao: modificacao,
	}

	c.JSON(http.StatusOK, gin.H{"data": completo})
}

func GetDisciplinasPodeCursar(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	var results []resultCodigo

	// pega o usuario
	var usuario, ok = controllers.GetLoginFromSession(c)
	if ok {
		query := db.Raw(queries.QUERY_PODE_CURSAR, sql.Named("usuario", usuario.CodUsuario))
		if err := query.Find(&results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": results})
}

func GetDisciplinasFaltaCursar(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	var results []resultCodigo

	// pega o usuario
	var usuario, ok = controllers.GetLoginFromSession(c)
	if ok {
		query := db.Raw(queries.QUERY_FALTA_CURSAR, sql.Named("usuario", usuario.CodUsuario))
		if err := query.Find(&results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": results})
}

// GetDisciplinaInfoCompleta retorna as informacoes completas de uma disciplina
//
// precisa retornar:
//
// - para uma disciplina:
// codigo, titulo, ementa, creditos,
// pre-requisitos, avaliacao, grau medio,
// quantidade de alunos que cursaram.
//
// - para uma turma:
// codigo, professor, shf, horario,
// destino, qtdVagas
func GetDisciplinaInfoCompleta(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	// pega o codigo da disciplina
	codDisciplina := c.Query("c")
	if codDisciplina == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Codigo da disciplina nao informado"})
		return
	}

	// pega as informacoes basicas da disciplina
	var resultsInfo []resultInfo
	query := db.Raw(queries.QUERY_DISCIPLINA_INFO, sql.Named("disciplina", codDisciplina))
	res := query.Find(&resultsInfo)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}
	// verifica se encontrou a disciplina
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Disciplina nao encontrada"})
		return
	}

	// pega as informacoes das turmas
	var resultsTurmas []resultTurmas
	query = db.Raw(queries.QUERY_TURMAS_INFO, sql.Named("disciplina", codDisciplina))
	if err := query.Find(&resultsTurmas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}

	// cria a resposta
	var resposta respostaDisciplina
	resposta.From(resultsInfo, resultsTurmas)

	c.JSON(http.StatusOK, gin.H{"data": resposta})
}

func GetDisciplinasCursadas(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	// pega o usuario
	var usuario, ok = controllers.GetLoginFromSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Usuario nao logado"})
		return
	}

	var results []models.Historico
	query := db.Where("cod_usuario = ?", usuario.CodUsuario)

	if err := query.Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}

	// converte os results para o formato de resposta
	// que é uma lista de codigos
	var resposta []resultCodigo
	for _, r := range results {
		resposta = append(resposta, resultCodigo{CodDisciplina: r.CodDisciplina})
	}

	c.JSON(http.StatusOK, gin.H{"data": resposta})
}
