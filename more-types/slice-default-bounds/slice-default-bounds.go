package main

import "fmt"

func main() {
	// Declare a slice literal of type []int
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	// Assign a smaller slice of the slice
	s = s[1:4]
	fmt.Println(s)

	// Assign a smaller slice, equal to s[0:2]
	s = s[:2]
	fmt.Println(s)

	// Assign a smaller slice, equal to s[1:s.length]
	s = s[1:]
	fmt.Println(s)

}
