package coach



type Coach struct {
	CoachId int `json:"coachId" example:"1"`
	FirstName string `json:"firstName" example:"Sergey"`
	LastName  string `json:"lastName" example:"Oleynicov"`
}