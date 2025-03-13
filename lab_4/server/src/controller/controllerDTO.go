package controller

type LoginRequest struct {
	Login    string `json:"login" example:"sergo_ranes"`
	Password string `json:"password" example:"password"`
}

type LoginResponse struct {
	CostumerId int    `json:"costumerId" example:"1"`
	FirstName  string `json:"firstName" example:"Sergey"`
	LastName   string `json:"lastName" example:"Medvedev"`
	Login      string `json:"login" example:"sergo_ranes"`
	Password   string `json:"password" example:"password"`
	Role       string `json:"access" example:"user"`
}

type RegistRequest struct {
	FirstName string `json:"firstName" example:"Sergey"`
	LastName  string `json:"lastName" example:"Medvedev"`
	Login     string `json:"login" example:"sergo_ranes"`
	Password  string `json:"password" example:"password"`
	Role      string `json:"access" example:"user"`
}

type CoachRecordResponse struct {
	FirstName string `json:"firstName" example:"Sergey"`
	LastName  string `json:"lastName" example:"Oleynicov"`
}



