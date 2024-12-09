package url

import (
	"lab_4/server/src/controller"
	
	"github.com/gin-gonic/gin"
)

func InitApiAuth(auth *gin.RouterGroup) {
	auth.POST("/login", controller.Login)
	auth.GET("/registration", controller.Registration)
	auth.POST("/registration-user",controller.RegistrationUser)
}


func InitApiSystem(system *gin.RouterGroup) {
	system.GET("/", controller.GetSystemPage)
	system.GET("/create-football-club", controller.GetPageCreateFootballClub)
	system.POST("/create-football-club-my-system", controller.CreateFootbalClubAndSquad)
	system.GET("/football-club", controller.GetFootballClubs)
	system.GET("/football-clubs", controller.GetFootballClub)
	system.GET("/squad", controller.GetPageSquad)
	system.GET("/squads", controller.GetSquad)
	system.GET("/coach", controller.GetPageCoach)
	system.GET("/footballer", controller.GetPageFootballer)
}

func InitApiCoach(coach *gin.RouterGroup) {
	coach.GET("/coaches", controller.GetCoaches)
}

func InitApiFootballer(coach *gin.RouterGroup) {
	coach.GET("/footballers", controller.GetFootballers)
}

