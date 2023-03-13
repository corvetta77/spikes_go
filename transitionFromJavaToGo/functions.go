package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	sub := func(x int, y int) int {
		return x - y
	}

	fmt.Println("add=", add(5, 4))
	fmt.Println("sub=", sub(5, 4))

	fmt.Printf("add(%v,%v) = %v\n", 5, 4, add(5, 4))
	fmt.Printf("sub(%v,%v) = %v\n", 5, 4, sub(5, 4))

}
