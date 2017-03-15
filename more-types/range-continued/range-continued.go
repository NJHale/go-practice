package main

import "fmt"

func main() {
	// Use make to instantiate a len(10) cap(10) slice with backing array of type [n]int
	pow := make([]int, 10)
	// For loop dropping the value from the tuple returned from range
	for i := range pow {
		pow[i] = 1 << uint(i) // 2 ** i multiply by 2 for each bit left-shifted
	}

	// Drop the first element of the return tuple each using '_'
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
