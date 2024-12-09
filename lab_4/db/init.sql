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
	football_club_id 	SERIAL       PRIMARY KEY,
	name_club CHARACTER VARYING(30)  NOT NULL,
	id_coach INT NOT NULL,
	CONSTRAINT fk_coach FOREIGN KEY(id_coach) REFERENCES coach(coach_id)
);

CREATE TABLE squad (
	squad_id		INT,
	name_squad	    CHARACTER VARYING(30)	NOT NULL,
	rating_squad	INT						NOT NULL,
	id_footballer	INT,
	id_football_club	INT,
	CONSTRAINT fk_footballer FOREIGN KEY(id_footballer) REFERENCES footballer(footballer_id),
	CONSTRAINT fk_football_club FOREIGN KEY(id_football_club) REFERENCES football_club(football_club_id) ON DELETE CASCADE
);

INSERT INTO coach (first_name, last_name) VALUES
('John', 'Smith'),
('Michael', 'Brown'),
('David', 'Johnson'),
('Robert', 'Williams'),
('James', 'Jones'),
('William', 'Garcia'),
('Thomas', 'Miller'),
('Daniel', 'Davis'),
('Paul', 'Martinez'),
('Mark', 'Hernandez');

INSERT INTO footballer (first_name, last_name, position, rating) VALUES
('Lionel', 'Messi', 'RW', 91), -- Нападающий (Правый фланг)
('Cristiano', 'Ronaldo', 'ST', 90), -- Нападающий (Центр)
('Robert', 'Lewandowski', 'ST', 91), -- Нападающий (Центр)
('Kylian', 'Mbappé', 'LW', 91), -- Нападающий (Левый фланг)
('Kevin', 'De Bruyne', 'CM', 91), -- Центральный полузащитник
('Neymar', 'Jr', 'LW', 89), -- Нападающий (Левый фланг)
('Mohamed', 'Salah', 'RW', 89), -- Нападающий (Правый фланг)
('Harry', 'Kane', 'ST', 89), -- Нападающий (Центр)
('Eden', 'Hazard', 'LW', 85), -- Нападающий (Левый фланг)
('Luka', 'Modrić', 'CM', 88), -- Центральный полузащитник
('Sergio', 'Ramos', 'CB', 86), -- Центральный защитник
('Virgil', 'van Dijk', 'CB', 89), -- Центральный защитник
('Thibaut', 'Courtois', 'GK', 90), -- Вратарь
('Jan', 'Oblak', 'GK', 89), -- Вратарь
('Toni', 'Kroos', 'CM', 87), -- Центральный полузащитник
('Zlatan', 'Ibrahimović', 'ST', 85), -- Нападающий (Центр)
('Sadio', 'Mané', 'LW', 89), -- Нападающий (Левый фланг)
('Antoine', 'Griezmann', 'CF', 87), -- Нападающий (Центр)
('Paul', 'Pogba', 'CM', 85), -- Центральный полузащитник
('Gerard', 'Piqué', 'CB', 85), -- Центральный защитник
('Raheem', 'Sterling', 'LW', 87), -- Нападающий (Левый фланг)
('Joshua', 'Kimmich', 'CDM', 88), -- Опорный полузащитник
('Bernardo', 'Silva', 'RW', 87), -- Нападающий (Правый фланг)
('Pierre-Emerick', 'Aubameyang', 'ST', 85), -- Нападающий (Центр)
('Frenkie', 'de Jong', 'CM', 87), -- Центральный полузащитник
('Jadon', 'Sancho', 'RW', 86), -- Нападающий (Правый фланг)
('David', 'De Gea', 'GK', 86), -- Вратарь
('Alisson', 'Becker', 'GK', 89), -- Вратарь
('Trent', 'Alexander-Arnold', 'RB', 87), -- Правый защитник
('Andrew', 'Robertson', 'LB', 87), -- Левый защитник
('Marc-André', 'ter Stegen', 'GK', 88), -- Вратарь
('Lautaro', 'Martínez', 'ST', 85), -- Нападающий (Центр)
('Mauro', 'Icardi', 'ST', 85), -- Нападающий (Центр)
('Hakim', 'Ziyech', 'RW', 84), -- Нападающий (Правый фланг)
('Gerard', 'Moreno', 'ST', 84), -- Нападающий (Центр)
('Sergio', 'Agüero', 'ST', 85), -- Нападающий (Центр)
('Ciro', 'Immobile', 'ST', 85), -- Нападающий (Центр)
('Christian', 'Pulisic', 'LW', 84), -- Нападающий (Левый фланг)
('Dusan', 'Tadic', 'LW', 85), -- Нападающий (Левый фланг)
('João', 'Felix', 'CF', 84), -- Нападающий (Центр)
('Mason', 'Mount', 'CM', 84), -- Центральный полузащитник
('Phil', 'Foden', 'LW', 85), -- Нападающий (Левый фланг)
('Angel', 'Di María', 'RW', 87), -- Нападающий (Правый фланг)
('Gareth', 'Bale', 'RW', 81), -- Нападающий (Правый фланг)
('James', 'Rodríguez', 'CAM', 84), -- Центральный атакующий полузащитник
('Mikel', 'Oyarzabal', 'LW', 84), -- Нападающий (Левый фланг)
('Rodri', 'Hernández', 'CDM', 86), -- Опорный полузащитник
('Kai', 'Havertz', 'CF', 85), -- Нападающий (Центр)
('Adama', 'Traoré', 'RW', 81), -- Нападающий (Правый фланг)
('Kingsley', 'Coman', 'LW', 85), -- Нападающий (Левый фланг)
('Fabinho', 'Silva', 'CDM', 85), -- Опорный полузащитник
('Nicolò', 'Barella', 'CM', 85), -- Центральный полузащитник
('Riyad', 'Mahrez', 'RW', 85); -- Нападающий (Правый фланг)
`