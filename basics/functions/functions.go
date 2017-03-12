package main

import "fmt"

func add(a int, b int) int {
  sum := a + b
  return sum
}

func main() {
  fmt.Println(add(43, 13))
}
