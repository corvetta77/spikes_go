package models

type City struct {
	Name  string
	TempC float64
}

// NewCity func that creates new City instance with given name
func NewCity(n string) *City {
	return &City{
		Name: n,
	}
}

// TempF returns the city temperature converted to Fahrenheit
func (c City) TempF() float64 {
	return (c.TempC * 9 / 5) + 32
}
