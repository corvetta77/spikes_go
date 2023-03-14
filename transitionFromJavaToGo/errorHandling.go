package main

import "fmt"

func area(x int, y int) (*int, error) {
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("zero area:[%v,%v]", x, y)
	}

	area := x * y
	return &area, nil
}

func main() {
	x := 3
	y := 1
	ar, err := area(x, y)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(ar)

	b, _ := area(2, 1)
	fmt.Printf("using naked return value to skip error handling. This line will print only if err not thrown: (%v,%v)=%v\n", x, y, *b)
}
