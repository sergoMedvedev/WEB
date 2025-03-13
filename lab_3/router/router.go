package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"lab_3/data"
	"log"
	"strconv"
)

func GetPlayer(c *gin.Context) {
	log.Println("/player/:id")
	ids := c.Param("id")
	log.Println(ids)
	idInt, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	for _, player := range data.Players {
		if player.Id == idInt {
			c.JSON(200, player)
			return
		}
	}
	c.JSON(404, gin.H{"error": "id not found"})
}

func CreatePlayer(c *gin.Context) {
	log.Println("/player")
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	pl := data.PlayerDTO{}

	err = json.Unmarshal(jsonData, &pl)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, player:= range data.Players {
		if player.Id == pl.Id {
			c.JSON(409, gin.H{"error": "dublicat id"})
			return
		}
	} 

	data.Players = append(data.Players, &pl)
	c.JSON(200, gin.H{"data": pl})
}

func DeletePlayer(c *gin.Context) {
	ids := c.Param("id")
	log.Println(ids)
	idInt, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	var clearPlayer []*data.PlayerDTO

	for _, plaer := range data.Players {
		if plaer.Id == idInt {
			continue
		}
		clearPlayer = append(clearPlayer, plaer)
	}

	data.Players = clearPlayer
	c.JSON(200, gin.H{
		"status": "OK",
		"data":   clearPlayer})
}

func GetPlayers(c *gin.Context) {
	c.JSON(200, gin.H{"data": data.Players})
}

func UpdatePlayer(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	pl := data.PlayerDTO{}

	err = json.Unmarshal(jsonData, &pl)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	var clearPlayer []*data.PlayerDTO

	for _, plaer := range data.Players {
		if plaer.Id == pl.Id {

			if plaer.Name != pl.Name {
				plaer.Name = pl.Name
			}

			if plaer.SecondName != plaer.SecondName {
				plaer.SecondName = pl.SecondName
			}

			if plaer.Rating != plaer.Rating {
				plaer.Rating = pl.Rating
			}

			clearPlayer = append(clearPlayer, plaer)
			continue
		}
		clearPlayer = append(clearPlayer, plaer)
	}
	c.JSON(200, gin.H{"data": pl})
}
