package model

type Player struct {
	id      string
	name    string
	surname string
	country string
}

func (p *Player) serName(name string) {
	p.name = name
}

func (p *Player) serSurname(surname string) {
	p.surname = surname
}

func (p *Player) serCountry(country string) {
	p.country = country
}
