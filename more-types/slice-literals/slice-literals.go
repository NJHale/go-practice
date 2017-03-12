package main

import "fmt"

func main() {
	// Declare a slice using a slice literal. This creates an underlying array
	// of type [n]int to reference as well as the slice of type []int
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// More of the same except with [n]bool and []bool
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// Declare a slice of type []struct{}, using an anonymous struct. This
	// creates an underlying array of type [n]struct{} to reference.
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
