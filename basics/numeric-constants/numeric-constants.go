package main

import "fmt"

// Declare large exported package level numeric constants
const (
	// Create a huge number by shifting a 1 bit left 100 places
	// In other words, the binary number 1 followed by 100 zeros
	Big = 1 << 100
	// Shift it right again by 99 places, so we end up with 1 << 1 = 2
	Small = Big >> 99
)

func needInt(x int64) int64 {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// The numeric constants are untyped and take the type needed by their context
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
