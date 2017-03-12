package main

import (
  "fmt"
  "math"
)

func main() {
  // Declare and initialize two ints
  var x, y int = 3, 4
  // Convert sum of the squares to float64
  var sumOfSquares float64 = float64(x*x + y*y)
  // Take the square root of the sum of the squares
  var sqrt float64 = math.Sqrt(sumOfSquares)
  // Convert to a uint
  var z uint = uint(sqrt);
  // Print out the variables
  fmt.Println(x, y, z)
}
