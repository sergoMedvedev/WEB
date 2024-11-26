package settings

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func InitializePostgres(db *sqlx.DB) {
	configUrl := "postgres:password@(localhost:15432)/test"
	db, err := sqlx.Connect("postgres", configUrl)
	if err != nil {
		log.Fatalln("Error when connect db: ", err.Error())
	}
}

func initPostgresDataBase() {

}
