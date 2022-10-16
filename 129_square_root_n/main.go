package main

import (
	"fmt"
	"math"
)

func solve(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	fmt.Printf("square root of %d is %.f\n", 9, solve(9))
}
