package main

import (
	"fmt"
	"github.com/Leinadium/puc-disciplinas/api/controllers"
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
	// base
	r := gin.Default()

	// para sessoes
	r.Use(sessions.Sessions("loginsession", cookie.NewStore([]byte(secretKey))))

	// configuracao do cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // temporario
		AllowCredentials: true,
	}))

	// rotas
	// rotas de autenticacao
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	r.GET("/login", controllers.CheckLogin)

	r.GET("/disciplinas/lista", controllers.GetDisciplinasLista)
	r.GET("/disciplinas/pesquisa", controllers.GetDisciplinasPesquisa)

	// rotas para a barra de pesquisa das disciplinas
	r.GET("/pesquisa/info", controllers.GetDisciplinasInformacoes)
	r.GET("/pesquisa/podecursar", controllers.GetDisciplinasPodeCursar)
	r.GET("/pesquisa/faltacursar", controllers.GetDisciplinasFaltaCursar)

	_ = r.Run()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Erro lendo .env file")
	}
}
