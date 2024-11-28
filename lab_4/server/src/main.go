package main

import (
	"lab_4/server/src/settings"

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
}
