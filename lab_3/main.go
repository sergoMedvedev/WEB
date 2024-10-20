package main

import (
	router "lab_3/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/player/{id}", router.getPlayer)


	r.Run("8800")
}
