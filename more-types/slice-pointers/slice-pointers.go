package main

import "fmt"

func main() {
	// Declare an array of type [n]string
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	// Declare two slices on the names array
	a := names[0:2] // Type of []string inferred
	b := names[1:3] // Type of []string inferred
	fmt.Println(a, b)

	// Change an element of the underlying array via a referencing slice
	b[0] = "XXX"
	// Show that changes are reflected in other slices referencing the names array
	// and the names array itself
	fmt.Println(a, b)
	fmt.Println(names)
}
