package main

import "fmt"

// this is one of those problems that doesn't come naturally to me
// i still struggle just thinking about the solution :(

var cache = map[int]int{}

func solve_child(arr []int, start int) int {
	if start >= len(arr) {
		return 0
	}

	curr := arr[start]
	max := 1

	count := 0

	for i := start + 1; i < len(arr); i++ {
		s := arr[i]
		if s >= curr {
			_, ok := cache[s]
			if ok {
				count = cache[s]
			} else {
				count = solve_child(arr, i) + 1
				cache[count] = count
			}

			if count > max {
				max = count
			}
		}
	}

	return max

}

func solve(arr []int) int {
	cache = map[int]int{}

	return solve_child(arr, 0)
}

func main() {
	arr := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	fmt.Println(arr, "==", solve(arr))
}
