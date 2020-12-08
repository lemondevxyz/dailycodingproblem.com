package main

import "fmt"

// constant space, variable time complexity
// ayyy
func solve(slice []int, size int) []int {
	if len(slice) < size {
		return nil
	}

	maxsum := 0
	ind := 0

	for i := 0; i < size; i++ {
		maxsum += slice[i]
	}

	for i := 1; i < len(slice); i++ {
		diff := len(slice) - i

		if diff >= size {
			bef := slice[i-1]
			cur := slice[i]

			val := maxsum + cur
			val = val - bef

			if val > maxsum {
				ind = i
				maxsum = val
			}
		}
	}

	if maxsum <= 0 {
		return nil
	}

	return slice[ind : ind+size]
}

func main() {
	inputs := [][]int{
		{34, -50, 42, 14, -5, 86},
		{-5, -1, -8, -9},
	}

	sizes := []int{4, 2}
	for k, v := range inputs {
		size := sizes[k]

		fmt.Printf("%v: %v\n", v, solve(v, size))
	}
}
