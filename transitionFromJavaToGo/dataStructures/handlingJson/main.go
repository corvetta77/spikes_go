package main

import (
	"fmt"
	"handlingJson/data"
	"handlingJson/models"
	"handlingJson/printer"
)

func main() {
	fmt.Println("Temperature Service")

	//create cities
	cities, err := models.NewCities(data.NewReader())
	if err != nil {
		fmt.Println("Fatal error ocurred")
		return
	}
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	for _, c := range cities.ListAll() {
		p.CityDetails(c)
	}
}
