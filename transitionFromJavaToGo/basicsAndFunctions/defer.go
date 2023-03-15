package main

import "fmt"

func printMsg(id string, m string) {
	fmt.Printf("[%v]: %v\n", id, m)
}

func main() {
	msg := "Hello, World"
	defer printMsg("Defer-0", msg)
	defer printMsg("Defer-1", msg)

	msg = "Goodbye"
	fmt.Println("Main has exited.")
}
