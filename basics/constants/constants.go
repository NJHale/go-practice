package main

import "fmt"

// Declare exported package level constant for pi
const Pi = 3.4 // Type inferred as float64

func main() {
	// Declare function level constant (not exported)
	const World = "世界" // Type inferred as string
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	// Declare another function level const (not exported)
	const Truth = true // Type inferred as bool
	fmt.Println("Go Rules?", Truth)
}
