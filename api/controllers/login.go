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

type microErrorProxy struct {
	Message string
	Code    int
}

type ResultCheck struct {
	Message   string `json:"message"`
	Usuario   string `json:"usuario"`
	Nome      string `json:"nome"`
	Curriculo bool   `json:"curriculo"`
}

// CheckLogin godoc
// @Summary Verifica o login do usuário
// @Description Testa o login do usuário através do cookie.
// @Produce json
// @Success 200 {object} ResultCheck "Sucesso"
// @Failure 401 {object} string "Não logado"
// @Router /login [get]
func CheckLogin(c *gin.Context) {
	usuario, ok := GetLoginFromSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Não logado",
		})
		return
	}

	// pra pegar o curriculo, precisa ver no banco
	var r ResultCheck
	r.Curriculo = false
	var db = GetDbOrSetError(c)
	if db != nil {
		// vendo no banco
		var u models.Usuario
		if err := db.Select("cod_curriculo").First(&u, "cod_usuario = ?", usuario.CodUsuario).Error; err == nil {
			r.Curriculo = u.CodCurriculo.Valid
		}
	}

	r.Message = "Logado"
	r.Usuario = usuario.CodUsuario
	r.Nome = usuario.NomeUsuario

	c.JSON(http.StatusOK, r)
}

// Login godoc
// @Summary Faz login
// @Description Faz login do usuário, setando o cookie
// @Param usuario formData string true "Usuário"
// @Param senha formData string true "Senha"
// @Produce json
// @Success 200 {object} string "Nome do usuário"
// @Failure 400 {object} string "Usuário ou senha não informados"
// @Failure 401 {object} string "Usuário ou senha incorretos"
// @Failure 500 {object} string "Erro ao acessar o serviço de login"
// @Router /login [post]
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
	session.Set(keyNomeSession, nome)
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

// Logout godoc
// @Summary Faz logout do usuario
// @Description Faz logout do usuario, resetando o cookie
// @Produce json
// @Success 200 {object} string "Logout realizado com sucesso"
// @Failure 500 {object} string "Erro ao salvar sessão"
// @Router /logout [post]
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

func fazLoginUsuarioMicro(usuario, senha string) (string, *microErrorProxy) {
	res, err := http.PostForm(urlMicroLogin, url.Values{
		"usuario": {usuario}, "senha": {senha},
	})
	if err != nil {
		return "", &microErrorProxy{Message: "Não foi possível acessar o serviço de login", Code: 500}
	}
	defer res.Body.Close()

	var bodyParsed map[string]string

	// primeiro, tenta fazer um match generico
	if err = json.NewDecoder(res.Body).Decode(&bodyParsed); err != nil {
		return "", &microErrorProxy{Message: "Responsa inválida do serviço de login", Code: 500}
	}

	// veja se o codigo foi erro e tem o campo de messagem na resposta
	if res.StatusCode >= 400 && bodyParsed["message"] != "" {
		return "", &microErrorProxy{Message: bodyParsed["message"], Code: res.StatusCode}
	}

	var nome = bodyParsed["nome"]
	if res.StatusCode == 200 && nome != "" {
		return nome, nil
	}

	// se nao deu match nem num erro, nem numa resposta, retorna erro
	return "", &microErrorProxy{Message: "Não foi possível fazer login", Code: 500}
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
