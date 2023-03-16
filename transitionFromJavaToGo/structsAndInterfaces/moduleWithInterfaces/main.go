package main

import (
	"fmt"
	"tempServiceWithInterface/models"
	"tempServiceWithInterface/printer"
)

func main() {
	fmt.Println("Temperature Service")

	//initialise printer and defer cleanup
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	//setup 3 cities
	lon := models.NewCity("London", 7.5)
	ams := models.NewCity("Amsterdam", 11)
	nyc := models.NewCity("New York", 3)

	p.CityDetails(lon)
	p.CityDetails(ams)
	p.CityDetails(nyc)

}
