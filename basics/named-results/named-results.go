package main

import "fmt"

func split(sum int) (x, y int) {
  // Split the sum in two
  x = sum * 3 / 9
  y = sum - x
  // Issue "naked" return statement
  return
}

func main() {
  fmt.Println(split(73))
}
