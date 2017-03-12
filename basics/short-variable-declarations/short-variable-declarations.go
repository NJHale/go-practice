package main

import "fmt"

func main() {
  // Use function level var declaration with explicit type and intializers
  var i, j int = 1, 2
  // Use var declaration short-hand with implicit type
  k := 3
  c, python, java := true, false, "no!"
  // Print out initialized variables
  fmt.Println(i, j, k, c, python, java)
}
