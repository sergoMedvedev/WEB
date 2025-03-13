package system

import (
	"lab_4/server/src/model/footballer"
	"lab_4/server/src/settings"
	"log"
	"time"
)

func CreateFootballClubAndSquad(request CreateFootbalClubAndSquadRequest) (CreateFootbalClubAndSquadRequest, error) {
	tx, err := settings.DBinst.Begin()
	if err != nil {
		log.Println("создание транзакции", err.Error())
		return CreateFootbalClubAndSquadRequest{}, err
	}
	// 1. Создание футбольного клуба
	var clubID int
	err = tx.QueryRow(`
		INSERT INTO football_club (name_club, id_coach)
		VALUES ($1, $2)
		RETURNING football_club_id`, request.ClubName, request.CoachID).Scan(&clubID)
	if err != nil {
		log.Println("ошибка при создании футбольного клуба", err.Error())
		return CreateFootbalClubAndSquadRequest{}, err
	}

	//найдем рейтинг состава
	arrayFottballer, err := footballer.GetFootballers()
	if err != nil {
		log.Println("ошибка при подсчете рейтинга состава", err.Error())
		return CreateFootbalClubAndSquadRequest{}, err
	}

	var rating int
	var sum int
	for _, footballer := range arrayFottballer {
		sum += footballer.Rating
	}

	rating = sum / len(arrayFottballer)

	for _, footballer_id := range request.FootballerIDs {
		squad_id := generateIdForSquad()

		
		_, err = tx.Exec(`
			INSERT INTO squad (squad_id, name_squad, rating_squad, id_footballer, id_football_club)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING squad_id`, squad_id, request.SquadName, rating, footballer_id, clubID)
		if err != nil {
			log.Println("ошика при создании состава", err.Error())
			return CreateFootbalClubAndSquadRequest{}, err
			} 
	}
	// Коммит транзакции
	err = tx.Commit()
	if err != nil {
		log.Println("Коммит транзакции")
		return CreateFootbalClubAndSquadRequest{}, err
	}
	return request, nil
}

func generateIdForSquad()int {
	timeU := time.Now().Unix()
	return int(timeU)
}

func GetFootballClub() ([]FootballClubRecordDTO, error) {
	footballClubList := make([]FootballClubRecordDTO, 0)

	rows, err := settings.DBinst.Queryx(`
		SELECT 
    	fc.football_club_id AS club_id,
    	fc.name_club AS club_name,
    	c.first_name AS coach_first_name,
		c.last_name AS coach_last_name,
    	s.name_squad AS squad_name,
    	s.rating_squad AS squad_rating
		FROM 
    	football_club fc
		LEFT JOIN 
    	coach c ON fc.id_coach = c.coach_id
		LEFT JOIN 
    	squad s ON fc.football_club_id = s.id_football_club;
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var footballId int = 0

	for rows.Next() {
		var footballClub FootballClubRecordDTO
		if err := rows.Scan(&footballClub.Club_id, &footballClub.Club_name, &footballClub.Coach_first_name, &footballClub.Coach_last_name, &footballClub.Squad_name, &footballClub.Squad_rating); 
		err != nil {
			return nil, err
		}
		if footballId == 0 {
			footballId = footballClub.Club_id
			footballClubList = append(footballClubList, footballClub)
		}

		if footballId != footballClub.Club_id {
			footballClubList = append(footballClubList, footballClub)
			footballId = footballClub.Club_id
		}
	}
	return footballClubList, nil
}


func GetSquadList() ([]SquadListRecordDTO, error) {
	squadList := make([]SquadListRecordDTO, 0)

	rows, err := settings.DBinst.Queryx(`
	SELECT 
	s.squad_id AS squad_id,
    fc.name_club AS club_name,
    s.name_squad AS squad_name,
    s.rating_squad AS squad_rating
	FROM 
    football_club fc
	LEFT JOIN 
    coach c ON fc.id_coach = c.coach_id
	LEFT JOIN 
    squad s ON fc.football_club_id = s.id_football_club
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var squadId int = 0

	for rows.Next() {
		var squad SquadListRecordDTO
		if err := rows.Scan(&squad.SquadId, &squad.ClubName, &squad.SquadName, &squad.SquadRating); 
		err != nil {
			return nil, err
		}
		if squadId == 0 {
			squadId = int(squad.SquadId)
			squadList = append(squadList, squad)
		}

		if squadId != int(squad.SquadId) {
			squadList = append(squadList, squad)
			squadId = int(squad.SquadId)
		}
	}
	return squadList, nil
}
