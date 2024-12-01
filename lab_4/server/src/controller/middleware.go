package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	role := session.Get("role")

	if role != "coach" && role != "player" && role != "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
		return
	}

	c.Next()
}
