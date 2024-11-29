package controller

import (
	"encoding/json"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

//____________AUTH________________________________

func Login(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	log.Println(string(jsonData))

	request := LoginRequest{}

	err = json.Unmarshal(jsonData, &request)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	log.Println(request)
}
