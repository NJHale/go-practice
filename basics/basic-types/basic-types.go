package main

import (
  "fmt"
  "math/cmplx"
)

// Factored package level var declaration
var (
  // Exported package level variables
  ToBe bool     = false
  MaxInt uint64 = 1<<64 - 1
  z complex128  = cmplx.Sqrt(-5 + 12i)
)

func main() {
  // Print the type and value of each variable
  fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T, Value %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T, Value: %v\n", z, z)
}
