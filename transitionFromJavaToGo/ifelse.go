package main

import "fmt"

func printParity_ifElse(x int) {
	if x%2 == 0 {
		fmt.Printf("%v is even.\n", x)
	} else {
		fmt.Printf("%v is odd.\n", x)
	}
}

func printParity_if(x int) {
	if x%2 == 0 {
		fmt.Printf("%v is even.\n", x)
		return
	}

	fmt.Printf("%v is odd.\n", x)
}

func printParity_shorthandVariable(x int) {
	if r := x % 2; r == 0 {
		fmt.Printf("%v is even.\n", x)
		return
	}

	fmt.Printf("%v is odd.\n", x)
}

func main6() {
	y := 1231
	x := 323232
	printParity_ifElse(y)
	printParity_ifElse(x)

	printParity_if(y)
	printParity_if(x)

	printParity_shorthandVariable(y)
	printParity_shorthandVariable(x)
}
