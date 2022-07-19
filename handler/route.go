package handler

import (
	"2022_07_HT/controller"

	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {

	e.Static("/img", "./img") // ./img에 접근을 허용

	e.POST("/signup", controller.SignUp)
	e.POST("/login", controller.Login)
	e.POST("/edituser", controller.EditUser)
	e.POST("/uploadimg", controller.UploadImg)
	e.POST("/uploadcontent", controller.UploadContent)

	e.GET("/viewstudent", controller.ViewStudent)
	e.GET("/viewgraduate", controller.ViewStudent)

}