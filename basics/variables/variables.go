package main

import "fmt"

// Make a var declaration at the package level
var c, python, java bool

func main() {
  // Make a var declaration at the function level
  var i int
  // Print out the default values of each variable
  fmt.Println(i, c, python, java)
}
