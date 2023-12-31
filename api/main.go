// @title 			puc-disciplinas API
// @version			1.0
// @description		API para o projeto Leinadium/puc-disciplinas
// @contact.name	Daniel
// @contact.url		https://github.com/Leinadium/puc-disciplinas
// @license.name	MIT
// @host			localhost:8080
// @BasePath		/

package main

import (
	"fmt"
	"os"

	"github.com/Leinadium/puc-disciplinas/api/controllers"
	"github.com/Leinadium/puc-disciplinas/api/controllers/avaliacoes"
	"github.com/Leinadium/puc-disciplinas/api/controllers/disciplinas"
	"github.com/Leinadium/puc-disciplinas/api/controllers/recomendacao"
	_ "github.com/Leinadium/puc-disciplinas/api/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// carrega as variaveis de ambiente do .env
	loadEnv()
	// pega a secret key
	secretKey, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		panic("SECRET_KEY nao encontrada")
	}
	// pega a url do front
	frontUrl, ok := os.LookupEnv("URL_FRONT")
	if !ok {
		frontUrl = "http://localhost:5173"
	}

	// construtor do gin
	r := gin.Default()

	// remove warning de trusted proxies
	r.SetTrustedProxies(nil)

	// para sessoes
	r.Use(sessions.Sessions("loginsession", cookie.NewStore([]byte(secretKey))))

	// configuracao do cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontUrl},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
	}))

	// para o envio de arquivos
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// para swagger
	// (steps para rodar o swagger):
	// ($ go install github.com/swaggo/swag/cmd/swag@latest && swag init --parseDependency)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// rotas
	// rotas de autenticacao
	authGroup := r.Group("")
	authGroup.POST("/login", controllers.Login)
	authGroup.POST("/logout", controllers.Logout)
	authGroup.GET("/login", controllers.CheckLogin)

	// rotas para a barra de pesquisa das disciplinas
	pesquisaGroup := r.Group("/pesquisa")
	pesquisaGroup.GET("/info", disciplinas.GetDisciplinasInformacoes)
	pesquisaGroup.GET("/podecursar", disciplinas.GetDisciplinasPodeCursar)
	pesquisaGroup.GET("/faltacursar", disciplinas.GetDisciplinasFaltaCursar)
	pesquisaGroup.GET("/cursadas", disciplinas.GetDisciplinasCursadas)

	// recomendacao
	recomendacaoGroup := r.Group("/recomendacao")
	recomendacaoGroup.POST("", recomendacao.GetRecomendacao)
	// informacao
	disciplinaGroup := r.Group("/disciplina")
	disciplinaGroup.GET("/info", disciplinas.GetDisciplinaInfoCompleta)
	// grade
	gradeGroup := r.Group("/grade")
	gradeGroup.POST("", controllers.PostGrade)
	gradeGroup.GET("", controllers.GetGrade)
	// historico
	historicoGroup := r.Group("/historico")
	historicoGroup.POST("", controllers.SubmitHistorico)
	// avaliacoes
	avaliacoesGroup := r.Group("/avaliacoes")
	avaliacoesGroup.GET("", avaliacoes.GetAvaliacoesTodos)
	avaliacoesGroup.POST("/disciplina", avaliacoes.PostAvaliacaoDisciplina)
	avaliacoesGroup.POST("/professor", avaliacoes.PostAvaliacaoProfessor)
	avaliacoesGroup.DELETE("/disciplina", avaliacoes.DeleteAvaliacaoDisciplina)
	avaliacoesGroup.DELETE("/professor", avaliacoes.DeleteAvaliacaoProfessor)
	// modificacao
	modificacaoGroup := r.Group("/modificacao")
	modificacaoGroup.GET("", controllers.GetModificacao)

	_ = r.Run()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Erro lendo .env file")
	}
}
