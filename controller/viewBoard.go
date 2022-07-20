package controller

import (
	db "2022_07_HT/database"
	"2022_07_HT/models"
	"log"

	"github.com/gin-gonic/gin"
)

type joinTable struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	Cardinal   string `json:"cardinal"`
	Num        int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT_NULL"`
	UserId     string `json:"userid" gorm:"NOT_NULL"`
	Title      string `json:"title" gorm:"NOT_NULL"`
	Content    string `json:"content" gorm:"NOT_NULL"`
}

func ViewBoard(c *gin.Context) {

	board := []models.Board{}
	user := []models.User{}

	err := db.DB.Find(&board).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "view board board finding error",
		})
		log.Println(err)
		return
	}

	err = db.DB.Find(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "view board board finding error",
		})
		log.Println(err)
		return
	}

	set := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(user); j++ {
			if board[i].UserId == user[j].Id {
				set++
			}
		}
	}

	joinedTable := make([]joinTable, set)

	set = 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(user); j++ {
			if board[i].UserId == user[j].Id {
				joinedTable[set].UserId = board[i].UserId
				joinedTable[set].Title = board[i].Title
				joinedTable[set].Content = board[i].Content
				joinedTable[set].Name = user[j].Name
				joinedTable[set].Department = user[j].Department
				joinedTable[set].Cardinal = user[j].Cardinal
				set++
			}
		}
	}

	c.JSON(200, gin.H{
		"message": "view board sunccess",
		"data":    joinedTable,
	})
}
