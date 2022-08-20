package main

import "fmt"

type Point struct {
	X int
	Y int
}

func CalculateSteps(arr []Point) int {
	if len(arr) == 0 {
		return 0
	}

	return len(arr) - 1
}

func solve(arr []Point) {
	fmt.Printf("Input: %v\nOutput: %d\n", arr, CalculateSteps(arr))
}

func main() {
	solve([]Point{{0, 0}, {1, 1}, {1, 2}})
}
