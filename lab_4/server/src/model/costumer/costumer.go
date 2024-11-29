package costumer


type Costumer struct {
	CostumerId int    `json:"costumerId" example:"1"`
	FirstName  string `json:"firstName" example:"Sergey"`
	LastName   string `json:"lastName" example:"Medvedev"`
	Login      string `json:"login" example:"sergo_ranes"`
	Password   string `json:"password" example:"password"`
	Role       string `json:"role" example:"user"`
}