package controller

import (
	"2022_07_HT/modules"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Upload IMG formFile Error",
		})
		return
	}

	fileName := header.Filename

	err = modules.MakeFile(file, fileName)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Upload IMG makeFile Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Upload IMG success",
	})

}
