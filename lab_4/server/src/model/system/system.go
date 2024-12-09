package system

type CreateFootbalClubAndSquadRequest struct {
	ClubName      string   `json:"club_name"`
	SquadName     string   `json:"squad_name"`
	CoachID       string   `json:"coach_id"`
	FootballerIDs []string `json:"footballer_ids"`
}

type FootballClubRecordDTO struct {
	Club_id          int    `json:"club_id"`
	Club_name        string `json:"club_name"`
	Coach_first_name string `json:"coach_first_name"`
	Coach_last_name  string `json:"coach_last_name"`
	Squad_name       string `json:"squad_name"`
	Squad_rating     string `json:"squad_rating"`
}

type SquadListRecordDTO struct {
	SquadId     int32  `json:"squad_id"`
	ClubName    string `json:"club_name"`
	SquadName   string `json:"squad_name"`
	SquadRating int    `json:"squad_rating"`
}
