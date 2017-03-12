package main

import "fmt"

func main() {
  // Declare typed variable and infer the type from the type of the rhs
  var i int
  h := i // int
  // Declare variables and infer the type from the numeric constant
  j := 2 + 8i // complex128
  k := 3.1415 // float64
  l := 21     // int
  // Print the types of all variables
  fmt.Printf("i: %T, h: %T, j: %T, k: %T, l: %T\n", i, h, j, k ,l)
}
