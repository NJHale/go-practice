package main

import "golang.org/x/tour/pic"



func Pic(dx, dy int) [][]uint8 {
	// Instantiate the return slice with make
	graph := make([][]uint8, dy)
	// Iterate through the graph slice, evaluate the equation, and store results
	for y := range graph {
		// Instantiate a slice of len and cap  dx stored at this y index with make
		graph[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			graph[y][x] = uint8(
				(x + y) / 2,
			)
		}
	}
  return graph
}

func main() {
	pic.Show(Pic)
}
