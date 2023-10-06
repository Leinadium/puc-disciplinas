package controllers

import (
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MicroErrorProxy struct {
	Message string
	Code    int
}

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
