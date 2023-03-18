package main

import "fmt"

func maxTempDestination(cityTemp map[string]float64) (string, float64) {
	var destWithMaxTemp string
	var maxTemp float64
	for city, temp := range cityTemp {
		if temp > maxTemp {
			maxTemp = temp
			destWithMaxTemp = city
		}
	}
	return destWithMaxTemp, maxTemp
}

func main() {
	temp := map[string]float64{
		"London": 20,
		"Krakow": 28,
		"Sydney": 32,
	}
	city, destination := maxTempDestination(temp)

	fmt.Printf("City '%v' has maxTemperature='%v'\n", city, destination)
}
