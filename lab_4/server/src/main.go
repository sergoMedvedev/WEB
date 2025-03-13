package main

import (
	"lab_4/server/src/controller"
	"lab_4/server/src/settings"
	"lab_4/server/src/url"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DBinst *sqlx.DB
var R *gin.Engine
var Store cookie.Store

func main() {

	//инициализация настроек
	settings.InitializeFromEnv()
	settings.InitializePostgres()
	log.Println(settings.Instance())

	R = gin.Default()

	Store = cookie.NewStore([]byte("your-secret-key"))
	R.Use(sessions.Sessions("mysession", Store))
	R.LoadHTMLGlob("template/*")
	R.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	authApi := R.Group("/auth")
	url.InitApiAuth(authApi)

	systemApi := R.Group("/my-system", controller.AuthMiddleware)
	url.InitApiSystem(systemApi)

	coachApi := R.Group("/coach", controller.AuthMiddleware)
	url.InitApiCoach(coachApi)

	footballerApi := R.Group("/footballer", controller.AuthMiddleware)
	url.InitApiFootballer(footballerApi)

	R.Run()
}
