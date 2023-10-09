package controllers

import (
	"database/sql"
	"github.com/Leinadium/puc-disciplinas/api/controllers/queries"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type resultLista struct {
	CodDisciplina  string `json:"codigo" gorm:"column:cod_disciplina"`
	NomeDisciplina string `json:"nome" gorm:"column:nome_disciplina"`
	Creditos       int    `json:"creditos" gorm:"column:creditos"`
	QtdVagas       int    `json:"qtdVagas" gorm:"column:qtdvagas"`
	QtdTurmas      int    `json:"qtdTurmas" gorm:"column:qtdturmas"`
}

func GetDisciplinasLista(c *gin.Context) {
	// pega o db
	var db = GetDbOrSetError(c)
	if db == nil {
		return
	}

	var results []resultLista

	// pega o usuario e executa a query
	var err error
	if usuario, ok := getLoginFromSession(c); ok {
		println("usuario: " + usuario)
		err = db.Raw(queries.QUERY_LISTA_LATERAL_LOGIN, sql.Named("name", usuario)).Scan(&results).Error
	} else {
		println("usuario nao logado")
		err = db.Raw(queries.QUERY_LISTA_LATERAL_ANON).Scan(&results).Error
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}

func GetDisciplinasPesquisa(c *gin.Context) {
	// pega o db
	var db = GetDbOrSetError(c)
	if db == nil {
		return
	}

	// trata o input
	var queryRaw = c.Query("query")
	strings.ReplaceAll(queryRaw, "%", "")
	strings.ReplaceAll(queryRaw, "_", "")
	strings.ReplaceAll(queryRaw, "'", "")
	strings.ReplaceAll(queryRaw, "\"", "")
	var query = "%" + c.Query("query") + "%"

	// faz a query
	var results []resultLista
	err := db.Raw(queries.QUERY_PESQUISA_LATERAL, sql.Named("query", query)).Scan(&results).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}
