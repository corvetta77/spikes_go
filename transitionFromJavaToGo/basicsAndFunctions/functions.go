package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main5() {
	sub := func(x int, y int) int {
		return x - y
	}

	fmt.Println("add=", add(5, 4))
	fmt.Println("sub=", sub(5, 4))

	fmt.Printf("add(%v,%v) = %v\n", 5, 4, add(5, 4))
	fmt.Printf("sub(%v,%v) = %v\n", 5, 4, sub(5, 4))

	x := 3

	negate := func(x *int) {
		neg := -*x
		*x = neg
	}

	fmt.Println(x)
	fmt.Println(&x)

	negate(&x)
	fmt.Println(x)
}
