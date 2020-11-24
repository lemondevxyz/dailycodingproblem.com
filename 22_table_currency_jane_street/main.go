// maybe i don't fully understand the solution, or the problem.
package main

import (
	"fmt"
	"math"
)

const maxint = int(^uint(0) >> 1)

type floats [][]float64

func (f floats) String() (str string) {

	for k, v := range f {
		str += fmt.Sprint(v)
		if k+1 < len(f) {
			str += "\n"
		}
	}

	return
}

func createGraph(rates [][]float64) {
	for k := range rates {
		arr := rates[k]
		for i := range arr {
			v := arr[i]
			rates[k][i] = -math.Log(v)
		}
	}
}

// thanks https://github.com/vaskoz/dailycodingproblem-go/blob/master/day32/problem.go
// really was lost to what the solution was.
// assume array is filled with positive integers
// O(3n^2) time, O(n) space
func solve(rates [][]float64, source int) bool {

	size := len(rates)
	createGraph(rates)
	minDist := make([]float64, size)
	for i := range minDist {
		minDist[i] = math.Inf(1)
	}
	minDist[source] = 0

	// turns out this loop is useless
	// so it's more O(n^2) than O(n^3)
	//for i := 0; i < len(rates)-1; i++ {
	for v := range rates {
		for w := range rates[v] {
			if minDist[w] > minDist[v]+rates[v][w] {
				minDist[w] = minDist[v] + rates[v][w]
			}
		}
	}
	// }

	for i := range rates {
		for j := range rates[i] {
			if minDist[j] > minDist[i]+rates[i][j] {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println("input:")
	rates := floats{
		{1, 0.741, 0.657, 1.061, 1.005},
		{1.349, 1, 0.888, 1.433, 1.366},
		{1.521, 1.126, 1, 1.614, 1.538},
		{0.942, 0.698, 0.619, 1, 0.953},
		{0.995, 0.732, 0.650, 1.049, 1},
	}
	fmt.Println(rates)
	fmt.Println("output:")
	fmt.Println(solve(rates, 0))

	fmt.Println("\ninput:")
	rates = [][]float64{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	fmt.Println(rates)
	fmt.Println("output")
	fmt.Println(solve(rates, 0))
}
