package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	// Create a key-value mapping for Bell Labs
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	// Create a key-value mapping for Google
	m["Google"] = Vertex{
		37.421999, -122.084057,
	}
  
	fmt.Println(m)
}
