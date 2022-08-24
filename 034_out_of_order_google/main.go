package main

// this problem has the solution written down
// i messed with this problem for like 5 minutes
// cool

import "fmt"

func solve(arr []int) int {
	inversion := 0
	compare := func(i, j int) {
		if arr[i] > arr[j] {
			if j > i {
				inversion++
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			compare(i, j)
		}
	}

	return inversion
}

func main() {
	inputs := [][]int{
		{2, 4, 1, 3, 5},
		{5, 4, 3, 2, 1},
	}

	for _, v := range inputs {
		fmt.Printf("%v: %d\n", v, solve(v))
	}
}
