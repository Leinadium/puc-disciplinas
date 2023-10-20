package main

import (
	"fmt"
	"github.com/Leinadium/puc-disciplinas/api/controllers"
	"github.com/Leinadium/puc-disciplinas/api/controllers/recomendacao"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// carrega as variaveis de ambiente do .env
	loadEnv()
	// pega a secret key
	secretKey, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		panic("SECRET_KEY nao encontrada")
	}

	// construtor do gin
	r := gin.Default()

	// para sessoes
	r.Use(sessions.Sessions("loginsession", cookie.NewStore([]byte(secretKey))))

	// configuracao do cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // temporario
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Type"},
	}))

	// rotas
	// rotas de autenticacao
	authGroup := r.Group("")
	authGroup.POST("/login", controllers.Login)
	authGroup.POST("/logout", controllers.Logout)
	authGroup.GET("/login", controllers.CheckLogin)
	// DEPRECIATED rotas para coletar as disciplinas
	disciplinasGroup := r.Group("/disciplinas")
	disciplinasGroup.GET("/lista", controllers.GetDisciplinasLista)
	disciplinasGroup.GET("/pesquisa", controllers.GetDisciplinasPesquisa)
	// rotas para a barra de pesquisa das disciplinas
	pesquisaGroup := r.Group("/pesquisa")
	pesquisaGroup.GET("/info", controllers.GetDisciplinasInformacoes)
	pesquisaGroup.GET("/podecursar", controllers.GetDisciplinasPodeCursar)
	pesquisaGroup.GET("/faltacursar", controllers.GetDisciplinasFaltaCursar)
	// recomendacao
	gradeGroup := r.Group("/grade")
	gradeGroup.POST("/recomendacao", recomendacao.GetRecomendacao)

	_ = r.Run()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Erro lendo .env file")
	}
}
