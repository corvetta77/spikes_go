package main

import "fmt"

var (
	v0 int
	v1 float64
	v2 bool
	v3 string
)

func main2() {
	fmt.Printf("v0 type=%T, value=%v \n", v0, v0)
	fmt.Printf("v1 type=%T, value=%v \n", v1, v1)
	fmt.Printf("v2 type=%T, value=%v \n", v2, v2)
	fmt.Printf("v3 type=%T, value=%v \n", v3, v3)
}
