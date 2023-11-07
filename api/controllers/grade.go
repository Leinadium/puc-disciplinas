package controllers

import (
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostGrade godoc
// @Summary Salva uma grade
// @Description Salva uma grade
// @Accept json
// @Produce json
// @Param grade body string true "Grade a ser salva"
// @Success 200 {object} string "id"
// @Failure 400 {object} string "Erro no corpo da requisicao"
// @Failure 401 {object} string "Nao logado"
// @Failure 500 {object} string "Erro ao salvar grade"
// @Router /grade [post]
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

// GetGrade godoc
// @Summary Coleta uma grade
// @Description Coleta uma grade
// @Produce json
// @Param id query string true "ID da grade"
// @Success 200 {object} string "data"
// @Failure 400 {object} string "ID nao fornecido"
// @Failure 404 {object} string "Grade nao encontrada"
// @Failure 500 {object} string "Erro ao conectar ao banco de dados"
// @Router /grade [get]
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
