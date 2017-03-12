package main

import "fmt"

func main() {
	fmt.Println("Stacking...")
	// Stack defers for lifo execution after main returns
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done stacking!")
}
