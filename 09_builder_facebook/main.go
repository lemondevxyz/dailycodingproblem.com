// i hate this problem :(
package main

import (
	"fmt"
	"os"
)

type min struct {
	min        int
	index      int
	second_min int
}

const maxint = int(^uint(0) >> 1)

// N is the number of houses
// K is the array of different colors
// we want to minimize cots while making sure that two neighboring houses are not of the same color
func solve(matrix [][]int) int {

	n := len(matrix)
	if n == 0 {
		return 0
	}

	k := len(matrix[0])

	prev := min{
		min:        0,
		index:      -1,
		second_min: 0,
	}

	for i := 0; i < n; i++ {
		curr := min{
			min:        maxint,
			index:      0,
			second_min: maxint,
		}

		for j := 0; j < k; j++ {
			if prev.index == j {
				matrix[i][j] += prev.second_min
			} else {
				matrix[i][j] += prev.min
			}

			if curr.min > matrix[i][j] {
				curr.second_min = curr.min
				curr.min = matrix[i][j]
				curr.index = j
			} else if curr.second_min > matrix[i][j] {
				curr.second_min = matrix[i][j]
			}
		}

		prev = curr
	}

	val := maxint
	for _, v := range matrix[n-1] {
		if val > v {
			val = v
		}
	}

	return val
}

var cases = map[int]int{
	solve([][]int{
		[]int{7, 3, 8, 6, 1, 2},
		[]int{5, 6, 7, 2, 4, 3},
		[]int{10, 1, 4, 9, 7, 6},
	}): 4,
	solve([][]int{
		[]int{7, 3, 8, 6, 1, 2},
		[]int{5, 6, 7, 2, 4, 3},
		[]int{10, 1, 4, 9, 7, 6},
		[]int{10, 1, 4, 9, 7, 6},
	}): 8,
}

func main() {

	for k, v := range cases {
		if k != v {
			fmt.Println("bad code")
			os.Exit(1)
		}
	}
	fmt.Println("success")

}
