package controller

import (
	db "2022_07_HT/database"
	"2022_07_HT/models"
	"log"

	"github.com/gin-gonic/gin"
)

func ViewBoard(c *gin.Context) {

	board := []models.Board{}

	err := db.DB.Find(&board).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "view board finding error",
		})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "view board sunccess",
		"data":    board,
	})
}
