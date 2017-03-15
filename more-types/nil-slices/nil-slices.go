package main

import "fmt"

func main() {
	// Declare a slice of type []int
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
