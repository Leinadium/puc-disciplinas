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

type InputProfessorPost struct {
	NomeProfessor string `json:"nome"`
	Nota          uint8  `json:"nota"`
}

type InputProfessorDelete struct {
	NomeProfessor string `json:"nome"`
}

type InputDisciplinaPost struct {
	CodDisciplina string `json:"codigo"`
	Nota          uint8  `json:"nota"`
}

type InputDisciplinaDelete struct {
	CodDisciplina string `json:"codigo"`
}

// GetAvaliacoesTodos godoc
// @Summary Coleta todas as avaliacoes
// @Description Coleta todas as avaliacoes
// @Produce json
// @Success 200 {object} avaliacoes.ResultGetAvaliacao "data"
// @Failure 401 {object} string "Nao logado"
// @Failure 500 {object} string "Erro ao executar query dos professores"
// @Failure 500 {object} string "Erro ao executar query das disciplinas"
// @Router /avaliacoes [get]
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
	var professores []ItemProfessorCompleto
	query := db.Raw(queries.QUERY_AVALIACOES_PROFESSOR, sql.Named("usuario", usuario.CodUsuario))
	if err := query.Find(&professores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query dos professores"})
		return
	}

	// pega as disciplinas
	var disciplinas []ItemDisciplinaCompleta
	query = db.Raw(queries.QUERY_AVALIACOES_DISCIPLINA, sql.Named("usuario", usuario.CodUsuario))
	if err := query.Find(&disciplinas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao executar query das disciplinas"})
		return
	}

	// retorna
	c.JSON(http.StatusOK, gin.H{
		"data": ResultGetAvaliacao{
			Professores: professores,
			Disciplinas: disciplinas,
		},
	})
}

// PostAvaliacaoProfessor godoc
// @Summary Salva uma avaliacao de professor
// @Description Salva uma avaliacao de professor
// @Accept json
// @Produce json
// @Param body body avaliacoes.InputProfessorPost true "body"
// @Success 200 {object} string "Avaliacao salva com sucesso"
// @Failure 400 {object} string "Requisicao invalida"
// @Failure 401 {object} string "Nao logado"
// @Failure 500 {object} string "Nao foi possivel salvar sua avaliacao"
// @Router /avaliacoes/professor [post]
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
	var input InputProfessorPost
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

// PostAvaliacaoDisciplina godoc
// @Summary Salva uma avaliacao de disciplina
// @Description Salva uma avaliacao de disciplina
// @Accept json
// @Produce json
// @Param body body avaliacoes.InputDisciplinaPost true "body"
// @Success 200 {object} string "Avaliacao salva com sucesso"
// @Failure 400 {object} string "Requisicao invalida"
// @Failure 401 {object} string "Nao logado"
// @Failure 500 {object} string "Nao foi possivel salvar sua avaliacao"
// @Router /avaliacoes/disciplina [post]
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
	var input InputDisciplinaPost
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

// DeleteAvaliacaoDisciplina godoc
// @Summary Deleta uma avaliacao de disciplina
// @Description Deleta uma avaliacao de disciplina
// @Accept json
// @Produce json
// @Param body body avaliacoes.InputDisciplinaDelete true "body"
// @Success 200 {object} string "Avaliacao deletada com sucesso"
// @Failure 400 {object} string "Requisicao invalida"
// @Failure 401 {object} string "Nao logado"
// @Failure 500 {object} string "Nao foi possivel deletar sua avaliacao"
// @Router /avaliacoes/disciplina [delete]
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
	var input InputDisciplinaDelete
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

// DeleteAvaliacaoProfessor godoc
// @Summary Deleta uma avaliacao de professor
// @Description Deleta uma avaliacao de professor
// @Accept json
// @Produce json
// @Param body body avaliacoes.InputProfessorDelete true "body"
// @Success 200 {object} string "Avaliacao deletada com sucesso"
// @Failure 400 {object} string "Requisicao invalida"
// @Failure 401 {object} string "Nao logado"
// @Failure 500 {object} string "Nao foi possivel deletar sua avaliacao"
// @Router /avaliacoes/professor [delete]
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
	var input InputProfessorDelete
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
