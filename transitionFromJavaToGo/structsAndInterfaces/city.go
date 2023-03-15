package main

import "fmt"

type city struct {
	name  string
	tempC float64
}

func newCity(n string) city {
	return city{
		name: n,
	}
}

func (c city) tempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

func (c *city) tempFModifyingState() float64 {
	return (c.tempC * 9 / 5) + 32
}

func tempF(c city) float64 {
	return (c.tempC * 9 / 5) + 32
}

func main() {
	c := city{
		name:  "London",
		tempC: 7.5,
	}
	tempF := tempF(c)
	fmt.Printf("[%v]: tempC=%v, tempF=%v\n", c.name, c.tempC, tempF)

	newCity := newCity("Krakow")
	newCity.tempC = 10
	tempFfromMethod := newCity.tempF()

	fmt.Printf("[%v]: tempC=%v, tempF=%v\n", newCity.name, newCity.tempC, tempFfromMethod)
}
