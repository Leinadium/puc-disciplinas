package controllers

import (
	"github.com/Leinadium/puc-disciplinas/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
