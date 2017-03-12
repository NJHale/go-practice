package main

import "fmt"

// Make the package level var declaration with initializers
var i, j int = 1, 2

func main() {
  // Make the package level var declaration with initializers, inferring types
  var c, python, java  = true, false, "no!"
  // Print the initialized variables
  fmt.Println(i, j, c, python, java)
}
