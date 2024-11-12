package model

type Coach struct {
	id      string
	name    string
	surname string
	country string
	squadId string
}

func (c *Coach) SetName(name string) {
	c.name = name
}

func (c *Coach) SetSurName(surname string) {
	c.surname = surname
}

func (c *Coach) SetCountry(country string) {
	c.country = country
}

func (c *Coach) GetId() string {
	return c.id
}

func (c *Coach) GetName() string {
	return c.name
}

func (c *Coach) GetSurname() string {
	return c.surname
}

func (c *Coach) GetCountry() string {
	return c.country
}
