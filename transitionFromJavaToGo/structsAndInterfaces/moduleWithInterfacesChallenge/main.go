package main

import (
	"fmt"
	"tempServiceWithInterfaceChallenge/models"
	"tempServiceWithInterfaceChallenge/printer"
)

func main() {
	fmt.Println("Temperature Service")

	//initialise printer and defer cleanup
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	//setup 3 cities
	lon := models.NewCity("London", 23, false, false)
	bar := models.NewCity("Barcelona", 30, true, false)
	nyc := models.NewCity("New York", 28, true, false)
	sta := models.NewCity("St. Anton", -3, false, true)
	asp := models.NewCity("Aspen", -2, false, true)

	p.CityDetails(lon)
	p.CityDetails(bar)
	p.CityDetails(nyc)
	p.CityDetails(sta)
	p.CityDetails(asp)

}
