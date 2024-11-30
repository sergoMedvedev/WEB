package url

import (
	"lab_4/server/src/controller"
	
	"github.com/gin-gonic/gin"
)

func InitApiAuth(auth *gin.RouterGroup) {
	auth.POST("/login", controller.Login)
	auth.GET("/registration", controller.Registration)
	auth.POST("/registration-user", controller.RegistrationUser)
}


func InitApiSystem(system *gin.RouterGroup) {
	system.GET("/", controller.GetSystemPage)
}

