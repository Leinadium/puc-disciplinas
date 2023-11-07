package recomendacao

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Leinadium/puc-disciplinas/api/controllers"
	"github.com/Leinadium/puc-disciplinas/api/controllers/queries"
	"github.com/gin-gonic/gin"
)

// GetRecomendacao godoc
// @Summary Coleta as disciplinas recomendadas
// @Description Coleta as disciplinas recomendadas
// @Produce json
// @Param body body recomendacao.InputRecomendacao true "Corpo da requisicao"
// @Success 200 {object} []ResultAlg "data"
// @Failure 400 {object} string "Erro ao ler corpo da requisicao"
// @Failure 500 {object} string "Erro ao conectar ao banco de dados"
// @Failure 500 {object} string "Erro ao executar query"
// @Router /recomendacao [post]
func GetRecomendacao(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	// verifica o login do usuario
	usuario, _ := controllers.GetLoginFromSession(c)

	// verifica o corpo da requisicao
	var input InputRecomendacao
	if err := json.NewDecoder(c.Request.Body).Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao ler corpo da requisicao"})
		return
	}

	// gerando a parte das turmas do sql
	// para cada escolha, é preciso gerar o fragmento sql:
	// (cod_disciplina = 'XXX0000' AND cod_turma = '0XX')
	// e junta-los com OR
	var escolhas string
	for i, escolha := range input.Escolhas {
		// fazendo algum tratamento no codigo e turma
		escolha.CodDisciplina = strings.ReplaceAll(escolha.CodDisciplina, "'", "")
		escolha.CodTurma = strings.ReplaceAll(escolha.CodTurma, "'", "")
		// colocando para maiusculo
		escolha.CodDisciplina = strings.ToUpper(escolha.CodDisciplina)
		escolha.CodTurma = strings.ToUpper(escolha.CodTurma)
		// criando query de escolhas
		escolhas += "(cod_disciplina = '" + escolha.CodDisciplina + "' AND cod_turma = '" + escolha.CodTurma + "')"
		if i < len(input.Escolhas)-1 {
			escolhas += " OR "
		}
	}
	// se nao tiver ainda nenhuma escolha, coloca 1 = 1 para retorna todas as escolhas
	if escolhas == "" {
		escolhas = "1 = 1"
	}

	// fazendo a substituicao das escolhas na query
	// OBS: caso o usuario nao esteja logado, o usuario sera ""
	//      o que faz com que o peso do conteudo sempre seja zero
	querySql := strings.Replace(queries.QUERY_RECOMENDACAO, "@escolhas", escolhas, 1)
	query := db.Raw(querySql, sql.Named("usuario", usuario.CodUsuario))
	// executando a query
	var results []resultQuery
	if err := query.Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}

	// executando o algoritmo
	resultsAlg := executaAlgoritmo(results)

	c.JSON(http.StatusOK, gin.H{"data": resultsAlg})
}
