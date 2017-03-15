package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// The second value in the return list from a retrieve is a bool representing
	// whether the given key is present in the map
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
