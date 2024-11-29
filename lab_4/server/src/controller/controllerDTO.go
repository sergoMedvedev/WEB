package controller


type Costumer struct {
	CostumerId int    `json:"costumerId" example:"1"`
	FirstName  string `json:"firstName" example:"Sergey"`
	LastName   string `json:"lastName" example:"Medvedev"`
	Login      string `json:"login" example:"sergo_ranes"`
	Password   string `json:"password" example:"password"`
	Access     string `json:"access" example:"user"`
}

type LoginRequest struct {
	Login      string `json:"login" example:"sergo_ranes"`
	Password   string `json:"password" example:"password"`
}

type LoginResponse struct {
	CostumerId int    `json:"costumerId" example:"1"`
	FirstName  string `json:"firstName" example:"Sergey"`
	LastName   string `json:"lastName" example:"Medvedev"`
	Login      string `json:"login" example:"sergo_ranes"`
	Password   string `json:"password" example:"password"`
	Access     string `json:"access" example:"user"`
}