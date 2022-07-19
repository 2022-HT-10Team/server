package controller

import (
	db "2022_07_HT/database"
	"2022_07_HT/helper"
	"2022_07_HT/models"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	user := &models.User{} // 회원 가입용 객체 생성

	err := c.Bind(user) // 프론트의 데이터 바인딩
	if err != nil {
		c.JSON(400, gin.H{
			"message": "login binding error",
		})
		return
	}

	err = db.DB.Where("id = ?", user.Id).Error // DB에 동일Id 존재 여부
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id already exist",
		})
		return
	}

	user.Pw, err = helper.HashPassword(user.Pw) // user pw 해싱
	if err != nil {
		c.JSON(400, gin.H{
			"message": "signUp pw hashing error",
		})
		return
	}

	err = db.DB.Create(user).Error // 정보를 변경한 유저정보를 db에 저장
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Create user Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "user Create success",
	})
}
