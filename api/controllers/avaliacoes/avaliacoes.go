package avaliacoes

import (
	"database/sql"
	"github.com/Leinadium/puc-disciplinas/api/controllers"
	"github.com/Leinadium/puc-disciplinas/api/controllers/queries"
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetAvaliacoesTodos(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	// pega o usuario
	var usuario, ok = controllers.GetLoginFromSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Não logado",
		})
		return
	}

	// pega os professores
	var professores []itemProfessorCompleto
	query := db.Raw(queries.QUERY_AVALIACOES_PROFESSOR, sql.Named("usuario", usuario.CodUsuario))
	if err := query.Find(&professores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query dos professores"})
		return
	}

	// pega as disciplinas
	var disciplinas []itemDisciplinaCompleta
	query = db.Raw(queries.QUERY_AVALIACOES_DISCIPLINA, sql.Named("usuario", usuario.CodUsuario))
	if err := query.Find(&disciplinas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query das disciplinas"})
		return
	}

	// retorna
	c.JSON(http.StatusOK, gin.H{
		"data": resultGetAvaliacao{
			Professores: professores,
			Disciplinas: disciplinas,
		},
	})
}

func PostAvaliacaoProfessor(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	// pega o usuario
	var usuario, ok = controllers.GetLoginFromSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "É necessário estar logado para avaliar",
		})
		return
	}

	// pega as informacoes
	var input struct {
		NomeProfessor string `json:"nome"`
		Nota          uint8  `json:"nota"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Requisição inválida"})
		return
	}

	// salva a avaliacao
	now := time.Now()
	obj := models.AvaliacaoProfessor{
		CodUsuario:    usuario.CodUsuario,
		NomeProfessor: input.NomeProfessor,
		NotaAvaliacao: input.Nota,
		DataAvaliacao: &now,
	}
	if err := db.Save(&obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível salvar sua avaliação"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Avaliacao salva com sucesso"})
	}
}

func PostAvaliacaoDisciplina(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	// pega o usuario
	var usuario, ok = controllers.GetLoginFromSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Não logado",
		})
		return
	}

	// pega as informacoes
	var input struct {
		CodDisciplina string `json:"codigo"`
		Nota          uint8  `json:"nota"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro no corpo da requisicao"})
		return
	}

	// salva a avaliacao
	now := time.Now()
	obj := models.AvaliacaoDisciplina{
		CodUsuario:    usuario.CodUsuario,
		CodDisciplina: input.CodDisciplina,
		NotaAvaliacao: input.Nota,
		DataAvaliacao: &now,
	}
	if err := db.Save(&obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao salvar avaliacao"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Avaliacao salva com sucesso"})
	}
}

func DeleteAvaliacaoDisciplina(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	// pega o usuario
	var usuario, ok = controllers.GetLoginFromSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Não logado",
		})
		return
	}

	// pega as informacoes
	var input struct {
		CodDisciplina string `json:"codigo"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro no corpo da requisicao"})
		return
	}

	// salva a avaliacao
	obj := models.AvaliacaoDisciplina{
		CodUsuario:    usuario.CodUsuario,
		CodDisciplina: input.CodDisciplina,
	}
	if err := db.Delete(&obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao deletar avaliacao"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Avaliacao deletada com sucesso"})
	}
}

func DeleteAvaliacaoProfessor(c *gin.Context) {
	// pega o db
	var db = controllers.GetDbOrSetError(c)
	if db == nil {
		return
	}

	// pega o usuario
	var usuario, ok = controllers.GetLoginFromSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Não logado",
		})
		return
	}

	// pega as informacoes
	var input struct {
		NomeProfessor string `json:"nome"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro no corpo da requisicao"})
		return
	}

	// salva a avaliacao
	obj := models.AvaliacaoProfessor{
		CodUsuario:    usuario.CodUsuario,
		NomeProfessor: input.NomeProfessor,
	}
	if err := db.Delete(&obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao deletar avaliacao"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Avaliacao deletada com sucesso"})
	}
}
