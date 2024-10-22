package router

import (
	"github.com/gin-gonic/gin"
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
		} else {
			c.JSON(404, gin.H{"error": "id not found"})
		}
	}
}
