package main

import "fmt"

func main3() {
	var ptr *string

	greeting := "Hello, world!"

	ptr = &greeting

	fmt.Println("Greeting:", greeting)
	fmt.Println("Address of greeting via reference:", &greeting)
	fmt.Println("Address of greeting via pointer name:", ptr)
	fmt.Println("Accessig value via pointer:", *ptr)

}
