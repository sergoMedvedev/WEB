package system

import (
	"fmt"
	"lab_4/server/src/model/footballer"
	"lab_4/server/src/settings"
	"log"
	"strings"
)

func CreateFootballClubAndSquad(request CreateFootbalClubAndSquadRequest) (CreateFootbalClubAndSquadRequest, error) {
	tx, err := settings.DBinst.Begin()
	if err != nil {
		log.Println("создание транзакции")
		return CreateFootbalClubAndSquadRequest{}, err
	}
	// 1. Создание футбольного клуба
	var clubID int
	err = tx.QueryRow(`
		INSERT INTO football_club (name_club, id_coach)
		VALUES ($1, $2)`, request.ClubName, request.CoachID).Scan(&clubID)
	if err != nil {
		log.Println("ошибка при создании футбольного клуба")
		return CreateFootbalClubAndSquadRequest{}, err
	}

	//найдем рейтинг состава
	arrayFottballer, err := footballer.GetFootballers()
	if err != nil {
		log.Println("ошибка при подсчете рейтинга состава")
		return CreateFootbalClubAndSquadRequest{}, err
	}

	var rating int 
	var sum int
	for _, footballer := range arrayFottballer {
		sum += footballer.Rating
	}

	rating = sum/len(arrayFottballer)


	// 2. Создание состава
	var squadID int
	err = tx.QueryRow(`
		INSERT INTO squad (name_squad, rating_squad, id_football_club)
		VALUES ($1, $2, $3)`, request.SquadName, rating, clubID).Scan(&squadID)
	if err != nil {
		log.Println("ошика при создании состава")
		return CreateFootbalClubAndSquadRequest{}, err
	}

	// 3. Привязка футболистов к составу
	// Формируем строку для IN запросов
	footballerIDs := make([]interface{}, len(request.FootballerIDs))
	for i, id := range request.FootballerIDs {
		footballerIDs[i] = id
	}
	placeholders := strings.Repeat("$", len(request.FootballerIDs))
	placeholders = strings.Join(strings.Split(placeholders, ""), ", ")
	query := fmt.Sprintf(`
		INSERT INTO squad_football_club (id_squad, id_football_club)
		SELECT $1, $2
		FROM unnest(%s) AS footballer_id;
	`, placeholders)

	_, err = tx.Exec(query, append([]interface{}{squadID, clubID}, footballerIDs...)...)
	if err != nil {
		log.Println("ошибка при ривязке футболистов к составу")
		return CreateFootbalClubAndSquadRequest{}, err
	}

	// Коммит транзакции
	err = tx.Commit()
	if err != nil {
		log.Println("Коммит транзакции")
		return CreateFootbalClubAndSquadRequest{}, err
	}
	return request, nil
}
