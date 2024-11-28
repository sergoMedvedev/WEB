package controller

import (
	"io"

	"github.com/gin-gonic/gin"
)

//____________AUTH________________________________

func Login(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
}
