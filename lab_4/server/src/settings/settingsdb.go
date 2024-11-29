package settings

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var DBinst *sqlx.DB
var initDB string = `
CREATE TABLE costumer (
	user_id 		SERIAL                PRIMARY KEY,
	first_name	    CHARACTER VARYING(30) NOT NULL,
	last_name 		CHARACTER VARYING(30) NOT NULL,
	login			CHARACTER VARYING(30) NOT NULL,
	password 		CHARACTER VARYING(30) NOT NULL,
	role			CHARACTER VARYING(30) NOT NULL
);

CREATE TABLE footballer (
	footballer_id 	SERIAL PRIMARY KEY,
	first_name	    CHARACTER VARYING(30) NOT NULL,
	last_name 		CHARACTER VARYING(30) NOT NULL,
	position 		CHARACTER VARYING(3)  NOT NULL,
	rating 			INT					  NOT NULL
);

CREATE TABLE coach (
	coach_id			SERIAL 				PRIMARY KEY,
	first_name	    CHARACTER VARYING(30) 	NOT NULL,
	last_name 		CHARACTER VARYING(30) 	NOT NULL
);

CREATE TABLE football_club (
	football_club_id 	SERIAL PRIMARY KEY,
	name_clab CHARACTER VARYING(30)  NOT NULL,
	id_coach INT NOT NULL,
	CONSTRAINT fk_coach FOREIGN KEY(id_coach) REFERENCES coach(coach_id)
);

CREATE TABLE squad (
	squad_id		SERIAL PRIMARY KEY,
	name_squad	CHARACTER VARYING(30)			NOT NULL,
	rating_squad	  INT						NOT NULL,
	id_footballer	  INT,
	id_football_club  INT,
	CONSTRAINT fk_footballer FOREIGN KEY(id_footballer) REFERENCES footballer(footballer_id),
	CONSTRAINT fk_football_club FOREIGN KEY(id_football_club) REFERENCES football_club(football_club_id)
);


CREATE TABLE squad_football_club (
	id_squad INT NOT NULL,
	id_football_club INT NOT NULL,
	CONSTRAINT fk_squad FOREIGN KEY(id_squad) REFERENCES squad(squad_id),
	CONSTRAINT fk_football_club FOREIGN KEY(id_football_club) REFERENCES football_club(football_club_id)
);
`

func InitializePostgres() {
	//configUrl := "postgresql://localhost:15432/test"
	db, err := sqlx.Connect("postgres", "user=postgres password=password dbname=test sslmode=disable")
	if err != nil {
		log.Fatalln("Error when connect db: ", err.Error())
	}
	DBinst = db
	log.Println("Successfully connected to postgres")
	initPostgresDataBase(db)
}

func initPostgresDataBase(db *sqlx.DB) {
	db.Exec(initDB)
}