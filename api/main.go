package main

import (
	"fmt"
	"github.com/Leinadium/puc-disciplinas/api/controllers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/disciplinas", controllers.GetDisciplinas)

	_ = r.Run()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Erro lendo .env file")
	}
}
