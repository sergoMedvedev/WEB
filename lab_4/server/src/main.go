package main

import (
	"lab_4/server/src/settings"
	"lab_4/server/src/url"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DBinst *sqlx.DB
var R *gin.Engine

func main() {

	//инициализация настроек
	settings.InitializeFromEnv()
	settings.InitializePostgres()

	R = gin.Default()

	R.LoadHTMLGlob("template/*")

	R.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	authApi := R.Group("/auth")
	url.InitApiAuth(authApi)

	systemApi := R.Group("/my-system")
	url.InitApiSystem(systemApi)

	R.Run()
}
