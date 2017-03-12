package main

import (
	"fmt"
)

func main() {
	// Defer the printing of "world" until the main function has returned
	defer fmt.Println("world")
	// Print "hello"
	fmt.Print("hello ")
}
