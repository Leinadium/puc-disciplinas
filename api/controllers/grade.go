package controllers

import (
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostGrade(c *gin.Context) {
	// pega o db
	var db = GetDbOrSetError(c)
	if db == nil {
		return
	}

	// pega o usuario
	var usuario, ok = GetLoginFromSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Nao logado"})
		return
	}

	// pega o corpo da requisicao
	input, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro no corpo da requisicao"})
		return
	}

	// salva o resultado, usando o models.Grade
	var grade = models.Grade{
		CodGrade:   GenerateRandomString(8),
		CodUsuario: usuario.CodUsuario,
		Conteudo:   string(input),
	}
	if err := db.Create(&grade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao salvar grade"})
		return
	}

	// retorna o codigo da grade
	c.JSON(http.StatusOK, gin.H{"id": grade.CodGrade})
}

func GetGrade(c *gin.Context) {
	// pega o db
	var db = GetDbOrSetError(c)
	if db == nil {
		return
	}

	// pega o id da grade
	var codGrade, ok = c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID não fornecido"})
		return
	}

	// procura no banco
	grade := models.Grade{CodGrade: codGrade}
	if err := db.First(&grade).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Grade não encontrada"})
		return
	}
	// retorna o conteudo
	c.JSON(http.StatusOK, gin.H{"data": grade.Conteudo})
}
