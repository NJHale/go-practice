package main

import "fmt"

func main() {
	// Instantiate array of prime numbers
	primes := [6]int{2, 3, 5, 7, 11, 13} // Type of [n]int is inferred on assignment
	// Declare a slice of type []int
	var s []int = primes[1:4]
	fmt.Println(s)
}
