package main

import "fmt"

func main() {
	// Instantiate a slice of type []int backed by a zeroed array of type [n]int
	// using make. Values are initialized to zero. a: len(5) cap(5)
	a := make([]int, 5)
	printSlice("a", a)

	// b: len(0) cap(5)
	b := make([]int, 0, 5)
  printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
