package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"lab_4/server/src/settings"
)

var DBinst *sqlx.DB

func main() {
	settings.InitializeFromEnv()
	settings.InitializePostgres(DBinst)
}
