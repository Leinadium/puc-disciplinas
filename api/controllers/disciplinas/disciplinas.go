package disciplinas

import (
	"database/sql"
	"github.com/Leinadium/puc-disciplinas/api/controllers"
	"github.com/Leinadium/puc-disciplinas/api/controllers/queries"
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetDisciplinasInformacoes godoc
// @Summary Coleta todas as disciplinas
// @Description Coleta as informações das disciplinas, além da data de modificação.
// @Produce json
// @Success 200 {object} []disciplinas.ResultLista "data"
// @Failure 500 {object} string "Erro ao conectar ao banco de dados"
// @Failure 500 {object} string "Erro ao executar query"
// @Router /pesquisa/info [get]
func GetDisciplinasInformacoes(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	// faz a query
	var results []ResultLista
	err := db.Raw(queries.QUERY_INFORMACOES).Scan(&results).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}

// GetDisciplinasPodeCursar godoc
// @Summary Coleta as disciplinas que o usuario pode cursar
// @Description Coleta o código das disciplinas que o usuário pode cursar
// @Produce json
// @Success 200 {object} []disciplinas.ResultCodigo "data"
// @Failure 500 {object} string "Erro ao conectar ao banco de dados"
// @Failure 500 {object} string "Erro ao executar query"
// @Router /pesquisa/podecursar [get]
func GetDisciplinasPodeCursar(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	var results []ResultCodigo

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

// GetDisciplinasFaltaCursar godoc
// @Summary Coleta as disciplinas que o usuario falta cursar
// @Description Coleta o código das disciplinas que o usuário falta cursar. É vazio se o usuário nao estiver logado.
// @Produce json
// @Success 200 {object} []disciplinas.ResultCodigo "data"
// @Failure 500 {object} string "Erro ao conectar ao banco de dados"
// @Failure 500 {object} string "Erro ao executar query"
// @Router /pesquisa/faltacursar [get]
func GetDisciplinasFaltaCursar(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	var results []ResultCodigo

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

// GetDisciplinaInfoCompleta godoc
// @Summary Coleta a informação completa de uma disciplina
// @Description Coleta a informação completa de uma disciplina
// @Produce json
// @Param c query string true "Codigo da disciplina"
// @Success 200 {object} []disciplinas.ResultInfo "data"
// @Failure 404 {object} string "Disciplina nao encontrada"
// @Failure 500 {object} string "Erro ao conectar ao banco de dados"
// @Failure 500 {object} string "Erro ao executar query"
// @Router /disciplina/info [get]
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
	var resultsInfo []ResultInfo
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
	var resultsTurmas []ResultTurmas
	query = db.Raw(queries.QUERY_TURMAS_INFO, sql.Named("disciplina", codDisciplina))
	if err := query.Find(&resultsTurmas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}

	// cria a resposta
	var resposta RespostaDisciplina
	resposta.From(resultsInfo, resultsTurmas)

	c.JSON(http.StatusOK, gin.H{"data": resposta})
}

// GetDisciplinasCursadas godoc
// @Summary Coleta os codigos das disciplinas que o usuario cursou
// @Description Coleta os codigos das disciplinas que o usuario cursou
// @Produce json
// @Success 200 {object} []disciplinas.ResultCodigo
// @Failure 401 {object} string "Usuario nao logado"
// @Failure 500 {object} string "Erro ao conectar ao banco de dados"
// @Failure 500 {object} string "Erro ao executar query"
// @Router /pesquisa/cursadas [get]
func GetDisciplinasCursadas(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	var results []models.Historico

	// pega o usuario
	var usuario, ok = controllers.GetLoginFromSession(c)
	if ok {
		query := db.Where("cod_usuario = ?", usuario.CodUsuario)
		if err := query.Find(&results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
			return
		}
	}

	// converte os results para o formato de resposta
	// que é uma lista de codigos
	resposta := make([]ResultCodigo, 0, len(results))

	for _, r := range results {
		resposta = append(resposta, ResultCodigo{CodDisciplina: r.CodDisciplina})
	}

	c.JSON(http.StatusOK, gin.H{"data": resposta})
}
