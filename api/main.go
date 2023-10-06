package main

import (
	"fmt"
	"github.com/Leinadium/puc-disciplinas/api/controllers"
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

	// rotas
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	r.GET("/disciplinas/lista", controllers.GetDisciplinasLista)

	_ = r.Run()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Erro lendo .env file")
	}
}
