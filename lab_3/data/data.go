package data

type Player struct {
	Id         int
	Name       string
	SecondName string
	Rating     int
}

var Players []*Player

func CreateTestdata() {
	Players = append(Players,
		&Player{Id: 1, Name: "Nikita", SecondName: "Orlov", Rating: 90},
		&Player{Id: 2, Name: "Kiril", SecondName: "Medvedev", Rating: 89},
		&Player{Id: 3, Name: "Oleg", SecondName: "Makarov", Rating: 70},
	)

}
