package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SubmitHistorico Rota de submissão de histórico
func SubmitHistorico(c *gin.Context) {
	// pega o usuario
	usuario, ok := GetLoginFromSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Não logado",
		})
		return
	}

	// le o arquivo
	var buf bytes.Buffer

	formFile, err := c.FormFile("historico")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao ler arquivo (1)",
		})
		return
	}

	openedFile, err := formFile.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao ler arquivo (2)",
		})
		return
	}
	defer openedFile.Close()
	file, err := io.ReadAll(openedFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao ler arquivo (3)",
		})
		return
	}

	// cria o campo do arquivo
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("historico", formFile.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao enviar histórico (1)",
		})
		return
	}
	_, err = part.Write(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao enviar histórico (2)",
		})
		return
	}
	writer.Close()

	// cria a requisicao
	finalUrl := urlMicroHistorico + "/?usuario=" + usuario.CodUsuario
	req, err := http.NewRequest("POST", finalUrl, &buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao enviar histórico (3)",
		})
		return
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// faz a requisicao
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao enviar histórico para o microserviço",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// decodificando o erro
		var bodyParsed map[string]string
		if err = json.NewDecoder(resp.Body).Decode(&bodyParsed); err != nil {
			// nao consigo nem decodificar o erro, entao erro generico
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao enviar histórico para o microserviço",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": bodyParsed["message"],
			})
		}
		return
	}

	type result struct {
		Inseridos int    `json:"inseridos"`
		Curriculo string `json:"curriculo"`
	}
	var bodyParsed result
	_ = json.NewDecoder(resp.Body).Decode(&bodyParsed)

	c.JSON(http.StatusOK, gin.H{
		"data": bodyParsed,
	})
}
