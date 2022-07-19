package controller

import (
	db "2022_07_HT/database"
	"2022_07_HT/models"

	"github.com/gin-gonic/gin"
)

func ViewMyBoard(c *gin.Context) {

	user := &[]models.User{}

	db.DB.Find(user)

	c.JSON(200, gin.H{
		"data": user,
	})
	db.DB.Preload("users.boards").First(&user)
}
