package main

import (
	"github.com/jmoiron/sqlx"
	"lab_4/server/src/settings"
)

var initSql string

var DBinst *sqlx.DB

func main() {
	settings.InitializeFromEnv()
	settings.InitializePostgres(DBinst)

}
