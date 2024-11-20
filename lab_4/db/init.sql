CREATE TABLE costumer (
	user_id 		SERIAL                PRIMARY KEY,
	first_name	    CHARACTER VARYING(30) NOT NULL,
	last_name 		CHARACTER VARYING(30) NOT NULL,
	login			CHARACTER VARYING(30) NOT NULL,
	password 		CHARACTER VARYING(30) NOT NULL,
	access			CHARACTER VARYING(30) NOT NULL
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