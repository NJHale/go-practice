package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len:%d cap:%d %v\n", len(s), cap(s), s)
}

func main() {
	// Instantiate slice of type []int using a slice literal. (An underlying array
	// is created to reference)
	s := []int{2, 3, 5, 7, 11, 13}

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// Reslice to extend its length
	s = s[:4]
	printSlice(s)

	// Reslice to drop the first two values
	s = s[2:]
	printSlice(s)

}
