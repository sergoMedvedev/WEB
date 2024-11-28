package controller

/*
user_id 		SERIAL                PRIMARY KEY,
	first_name	    CHARACTER VARYING(30) NOT NULL,
	last_name 		CHARACTER VARYING(30) NOT NULL,
	login			CHARACTER VARYING(30) NOT NULL,
	password 		CHARACTER VARYING(30) NOT NULL,
	access			CHARACTER VARYING(30) NOT NULL
*/

type Costumer struct {
	CostumerId int
	firstName string
	lastName string
	login string
	password string
	access string
} 

