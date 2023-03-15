package main

import "fmt"

func rectangle(x int, y int) (int, int) {
	area := x * y
	circumf := 2 * (x + y)
	return area, circumf
}

func rectangleWithNames(x int, y int) (area int, circumf int) {
	area = x * y
	circumf = 2 * (x + y)
	return
}

func main7() {
	x := 2
	y := 3

	a, c := rectangle(x, y)
	fmt.Printf("Rectangle of (%v, %v), area = %v;\n", x, y, a)
	fmt.Printf("Rectangle of (%v, %v), circumference = %v;\n", x, y, c)

	a2, c2 := rectangleWithNames(x, y)
	fmt.Printf("Rectangle of (%v, %v), area = %v;\n", x, y, a2)
	fmt.Printf("Rectangle of (%v, %v), circumference = %v;\n", x, y, c2)

	//using blank identifier
	a3, _ := rectangleWithNames(x, y)
	fmt.Printf("Rectangle of (%v, %v), area = %v;\n", x, y, a3)
	// fmt.Printf("Rectangle of (%v, %v), circumference = %v;\n", x, y, c2)
}
