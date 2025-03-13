package footballer

type Footballer struct {
	FootballerId int    `json:"footballerId" example:"1"`
	FirstName    string `json:"firstName" example:"Sergey"`
	LastNane     string `json:"lastName" example:"Roben"`
	Position     string `json:"position" example:"HAF"`
	Rating       int    `json:"rating" example:"80"`
}
