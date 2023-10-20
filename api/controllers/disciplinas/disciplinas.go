package disciplinas

import (
	"database/sql"
	"github.com/Leinadium/puc-disciplinas/api/controllers"
	"github.com/Leinadium/puc-disciplinas/api/controllers/queries"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	c.JSON(http.StatusOK, gin.H{"data": results})
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
		println("usuario: " + usuario)
		// query := db.Model(&models.Historico{}).Select("cod_disciplina").Where("cod_usuario = ?", usuario)
		query := db.Raw(queries.QUERY_PODE_CURSAR, sql.Named("name", usuario))
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
		query := db.Raw(queries.QUERY_FALTA_CURSAR, sql.Named("name", usuario))
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
	codDisciplina := c.Query("codigo")
	if codDisciplina == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Codigo da disciplina nao informado"})
		return
	}

	// pega as informacoes basicas da disciplina
	var resultsInfo []resultInfo
	query := db.Raw(queries.QUERY_DISCIPLINA_INFO, sql.Named("disciplina", codDisciplina))
	if err := query.Find(&resultsInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}

	// pega as informacoes das turmas
	var resultsTurmas []resultTurmas
	query = db.Raw(queries.QUERY_TURMAS_INFO, sql.Named("disciplina", codDisciplina))
	if err := query.Find(&resultsTurmas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}
}
