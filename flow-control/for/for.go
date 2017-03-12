package main

import "fmt"

func main() {
	sum := 0
  // Sum of i from [0, 9]
	for i := 0; i < 10; i++ {
		sum += i
	}
  // General solution -> n(n+1)/2
	fmt.Println(sum)
}
