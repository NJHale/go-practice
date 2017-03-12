package main

import "fmt"

func main() {
  // Declare a string array of size 2
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)
  // Declare and instantiate array of 6 ints
  primes := [6]int{2, 3, 5, 7, 11, 13} // Type [n]int is inferred on assignment
  fmt.Println(primes)
}
