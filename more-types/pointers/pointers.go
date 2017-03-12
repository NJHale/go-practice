package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // Point to i
	fmt.Println(*p) // Read i through pointer (by dereferencing pointer p)
	*p = 21         // Set i through pointer (by dereferencing pointer p)
	fmt.Println(i)  // See the new value of i

	p = &j         // Point to j
	*p = *p / 37   // Divide through pointer (by dereferencing pointer p)
	fmt.Println(j) // See the new value of j
}
