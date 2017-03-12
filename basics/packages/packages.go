// Specify package main to indicate the program's entrypoint
package main

import (
  "fmt"
  "math/rand"
)

func main() {
  fmt.Println("My favorite number is ", rand.Intn(10))
}
