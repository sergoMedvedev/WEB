package main

import (
	"lab_4/server/src/settings"

	"github.com/jmoiron/sqlx"
)

var initSql string

type client struct {
	dbinst *sqlx.DB
}

func main() {
	settings.InitializeFromEnv()
	

}
