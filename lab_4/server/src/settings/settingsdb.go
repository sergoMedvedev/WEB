package settings

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func InitializePostgres(db *sqlx.DB) {
	//configUrl := "postgresql://localhost:15432/test"
	db, err := sqlx.Open("postgres", "user=postgres password=password dbname=test sslmode=disable")
	if err != nil {
		log.Fatalln("Error when connect db: ", err.Error())
	}
	log.Println("Successfully connected to postgres")

}

func initPostgresDataBase(db *sqlx.DB) {
	
}
