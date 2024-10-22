package main

import (
	"github.com/gin-gonic/gin"
	"lab_3/data"
	"lab_3/router"
)

func main() {
	//Определим тестовые данные
	data.CreateTestdata()

	r := gin.Default()

	r.GET("/player/:id", router.GetPlayer)
	r.POST("/player", router.CreatePlayer)
	r.DELETE("/player/:id", router.DeletePlayer)
	r.GET("/players", router.GetPlayers)
	r.PUT("/player", router.UpdatePlayer)

	r.Run()
}
