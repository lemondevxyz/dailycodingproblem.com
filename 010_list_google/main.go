// good problem
package main

import (
	"fmt"
	"os"
)

func solve(prev, next []int) int {

	m := len(prev)
	n := len(next)
	if n > m {
		n = m
	}

	for i := 0; i < n; i++ {
		if prev[i] == next[i] {
			return prev[i]
		}
	}

	// bad list
	return -1
}

var cases = map[int]int{
	solve([]int{3, 7, 8, 10}, []int{99, 1, 8, 10}): 8,
	solve([]int{6, 9, 1, 9}, []int{65, 48, 24, 9}): 9,
	solve([]int{4, 2, 1}, []int{6, 2, 9, 5}):       2,
	solve([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}):    -1,
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
