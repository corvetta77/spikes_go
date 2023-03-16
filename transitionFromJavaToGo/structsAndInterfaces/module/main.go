package main

import (
	"fmt"
	"tempService/models"
	"tempService/printer"
)

func main() {
	fmt.Println("Temperature Service")

	//initialise printer and defer cleanup
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	//setup 3 cities
	lon := models.NewCity("London")
	lon.TempC = 7.5

	ams := models.NewCity("Amsterdam")
	ams.TempC = 11

	nyc := models.NewCity("New York")
	nyc.TempC = -3

	p.CityDetails(lon)
	p.CityDetails(ams)
	p.CityDetails(nyc)

}
