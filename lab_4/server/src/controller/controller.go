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

	c.Set("id", user.CostumerId)
	c.Set("role", user.Role)

	c.JSON(200, "{status : ok}")
}

func Registration(c *gin.Context) {
	c.JSON(200, gin.H{"status": "OK"})
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

	c.JSON(200, "")
}

//____________SYSTEM________________________________

func GetSystemPage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}


