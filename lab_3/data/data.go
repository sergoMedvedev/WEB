package data

type Player struct{
	Id int
	Name string
	SecondName string
	Rating int
}

player1 := player{
	Id: 1,
	Name: "Nikita",
	SecondName: "Medvedev",
	Rating: 90,
}

var players []*Player = []*Player{
	&player1,
}