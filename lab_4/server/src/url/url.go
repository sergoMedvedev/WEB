package url

import (
	"lab_4/server/src/controller"
	
	"github.com/gin-gonic/gin"
)

func InitApiAuth(auth *gin.RouterGroup) {
	auth.POST("/login", controller.Login)
}