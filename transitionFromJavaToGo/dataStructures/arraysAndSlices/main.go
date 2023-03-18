package main

import "fmt"

func addCity(city string, cities *[2]string) {
	cities[1] = city
}

func main() {
	cities := [2]string{"London"}
	copy := cities

	cities[1] = "New York"
	addCity("Miami", &copy)

	fmt.Println("Cities: ", cities)
	fmt.Println("Copies: ", copy)

	citiesSlice := make([]string, 1, 1)
	copySlice := citiesSlice
	citiesSlice[0] = "London"
	citiesSlice = append(citiesSlice, "New York")

	fmt.Printf("[%p]cities=%v, len = %v, cap=%v \n", &citiesSlice, citiesSlice, len(citiesSlice), cap(citiesSlice))
	fmt.Printf("[%p]copy=%v, len = %v, cap=%v \n", &copySlice, copySlice, len(copySlice), cap(copySlice))

}
