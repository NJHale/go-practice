package main

import "fmt"

func main() {
  // Make package level var declaration with no initialization
  var i int
  var f float64
  var b bool
  var s string
  // Print out the value of each variable
  fmt.Printf("%v, %v, %v, %q\n", i, f, b, s)
}
