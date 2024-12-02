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
