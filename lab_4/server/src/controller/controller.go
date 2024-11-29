package controller

import (
	"encoding/json"
	"io"
	"log"

	"lab_4/server/src/model/costumer"

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

	request := costumer.Costumer{}

	err = json.Unmarshal(jsonData, &request)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	user, err := costumer.LoginCostumer(c, &request)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid login or password"})
	}

	log.Println(user.CostumerId)
	log.Println(user.Role)
	c.Set("id", user.CostumerId)
	c.Set("role", user.Role)

	c.HTML(200, "index.html", nil)
}

func Registration(c *gin.Context) {
	c.HTML(200, "registration.html", nil)
}

func RegistrationUser(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	log.Println(string(jsonData))

	request := costumer.Costumer{}

	err = json.Unmarshal(jsonData, &request)
	if err != nil {
		log.Println("400 блллл")
		c.JSON(400, err.Error())
		return
	}

	err = costumer.RegistCostumer(c, &request)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error server"})
	}

	c.HTML(200, "index.html", nil)
}


