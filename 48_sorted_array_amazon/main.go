package main

import "fmt"

// cool
func solve(arr []int, element int) int {
	// given that it's been rotated, if loop value is bigger than element then just go back...
	for i := 0; i < len(arr); {
		v := arr[i]

		if v > element {
			i--
		} else if v < element {
			i++
		} else if v == element {
			return i
		}

		if i == -1 {
			i = len(arr) - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(solve([]int{13, 18, 25, 2, 8, 10}, 8))
}
