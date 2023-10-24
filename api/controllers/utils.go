package controllers

import (
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

func GetDbOrSetError(c *gin.Context) *gorm.DB {
	db, err := models.GetDB()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Erro ao conectar ao banco de dados",
		})
		return nil
	}
	return db
}

// GetLoginFromSession Pega o login a partir da sessao
// Retorna o nome do usuario ou "" se nao estiver logado
func GetLoginFromSession(c *gin.Context) (string, bool) {
	session := sessions.Default(c)
	usuario := session.Get(keyLoginSession)
	if usuario == nil {
		return "", false
	}
	return usuario.(string), true
}

type RoundedFloat float64

func (f RoundedFloat) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(f)*100, 'f', 1, 32)), nil
}

var letras = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(n int) string {
	// gera uma seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// gera um sequencia
	b := make([]rune, n)
	for i := range b {
		b[i] = letras[r.Intn(len(letras))]
	}
	return string(b)
}
