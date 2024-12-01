package system

type CreateFootbalClubAndSquadRequest struct {
	ClubName      string   `json:"club_name"`
	SquadName     string   `json:"squad_name"`
	CoachID       string   `json:"coach_id"`
	FootballerIDs []string `json:"footballer_ids"`
}