package data

type PlayerDTO struct {
	Id         int
	Name       string
	SecondName string
	Rating     int
}

var Players []*PlayerDTO

func CreateTestdata() {
	Players = append(Players,
		&PlayerDTO{Id: 1, Name: "Nikita", SecondName: "Orlov", Rating: 90},
		&PlayerDTO{Id: 2, Name: "Kiril", SecondName: "Medvedev", Rating: 89},
		&PlayerDTO{Id: 3, Name: "Oleg", SecondName: "Makarov", Rating: 70},
	)

}
