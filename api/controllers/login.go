package controllers

import (
	"encoding/json"
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
	"net/url"
)

// CheckLogin Rota de checagem de login
func CheckLogin(c *gin.Context) {
	usuario, ok := GetLoginFromSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Não logado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logado",
		"nome":    usuario,
	})
}

// Login Rota de Login
func Login(c *gin.Context) {
	usuario := c.PostForm("usuario")
	senha := c.PostForm("senha")

	if usuario == "" || senha == "" {
		c.JSON(400, gin.H{
			"message": "Usuário ou senha não informados",
		})
		return
	}

	// repassa para o serviço de login
	nome, err := fazLoginUsuarioMicro(usuario, senha)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"message": err.Message,
		})
		return
	}

	// checa se o usuario ja existe ou precisa criar
	checkNovoUsuario(c, usuario, nome)

	// adiciona na sessao
	session := sessions.Default(c)
	session.Set(keyLoginSession, usuario)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao salvar sessão",
		})
		return
	}

	// retorna o nome dele
	c.JSON(http.StatusOK, gin.H{
		"nome": nome,
	})
}

// Logout Rota de Logout
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao salvar sessão",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout realizado com sucesso",
	})
}

func fazLoginUsuarioMicro(usuario, senha string) (string, *MicroErrorProxy) {
	res, err := http.PostForm(urlMicroLogin, url.Values{
		"usuario": {usuario}, "senha": {senha},
	})
	if err != nil {
		return "", &MicroErrorProxy{Message: "Não foi possível acessar o serviço de login", Code: 500}
	}
	defer res.Body.Close()

	var bodyParsed map[string]string

	// primeiro, tenta fazer um match generico
	if err = json.NewDecoder(res.Body).Decode(&bodyParsed); err != nil {
		return "", &MicroErrorProxy{Message: "Responsa inválida do serviço de login", Code: 500}
	}

	// veja se o codigo foi erro e tem o campo de messagem na resposta
	if res.StatusCode >= 400 && bodyParsed["message"] != "" {
		return "", &MicroErrorProxy{Message: bodyParsed["message"], Code: res.StatusCode}
	}

	var nome = bodyParsed["nome"]
	if res.StatusCode == 200 && nome != "" {
		return nome, nil
	}

	// se nao deu match nem num erro, nem numa resposta, retorna erro
	return "", &MicroErrorProxy{Message: "Não foi possível fazer login", Code: 500}
}

// checkNovoUsuario checa se o usuario ja existe ou precisa criar
func checkNovoUsuario(c *gin.Context, usuario, nome string) {
	var db = GetDbOrSetError(c)
	if db == nil {
		return
	}

	// vai dar erro se o usuario ja existir
	// mas tudo bem
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.Usuario{
		CodUsuario: usuario, NomeUsuario: nome,
	})

}
