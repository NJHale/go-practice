package main

import (
	"fmt"
	"math"
)

func newtonsMethod(z, x, epsilon float64) (float64, int64) {
	// z_n+1 = z_n  - (z_n^2 - x) / 2z_n
	// Iterative solution
	// Declare counter
	var n int64
	// Perform first iteration outside loop
	zNext := z - (z*z-x)/2*z
	// Loop while change in z >= epsilon
	for math.Abs(zNext-z) >= epsilon {
		z = zNext
		zNext = z - (z*z-x)/2*z
		//fmt.Println("zNext:", zNext)
		// Increment the counter
		n++
	}
	return zNext, n
}

func Sqrt(x float64) float64 {
	// Use newtonsMethod to calculate
	est, n := newtonsMethod(.9, x, 0.0001)
	fmt.Println("Number of Newtons Method iterations:", n)
	return est
}

func main() {
	fmt.Printf("Sqrt: %g, math.Sqrt: %g\n", Sqrt(2), math.Sqrt(2))
}
