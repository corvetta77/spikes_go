package main

import "fmt"

func main() {
	var ptr *string

	greeting := "Hello, world!"

	ptr = &greeting

	fmt.Println("Greeting:", greeting)
	fmt.Println("Address of greeting:", &greeting)
	fmt.Println("Address of greeting:", ptr)
	fmt.Println("Accessig value via pointer:", *ptr)

}
