package controllers

import (
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-gonic/gin"
)

// GetModificacao godoc
// @Summary Coleta as datas da ultima modificacao
// @Description Coleta as datas da ultima modificacao
// @Produce json
// @Success 200 {object} models.Modificacao "data"
// @Failure 500 {object} string "Erro ao conectar ao banco de dados"
// @Router /modificacao [get]
func GetModificacao(c *gin.Context) {
	// pega o db
	var db = GetDbOrSetError(c)
	if db == nil {
		return
	}

	// pega as modificacoes (se der erro, tudo bem)
	var modificacao models.Modificacao
	db.First(&modificacao)

	// retorna a modificacao
	c.JSON(200, gin.H{"data": modificacao})
}
