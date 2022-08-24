// ngl my approach was pretty inefficient.
// turns out using dynamic programming + a bottom ups manner is way better than two nested loops
package main

import (
	"fmt"
	"os"
)

// n = steps in staircase
// x = allowed steps
func solve(n int, x []int) int {
	if n == 0 {
		return 1
	}

	nums := make([]int, n+1)
	nums[0] = 1

	for i := 1; i <= n; i++ {
		total := 0
		for _, v := range x {
			if i-v >= 0 {
				total += solve(i-v, x)
			}
		}

		nums[i] = total
	}

	return nums[n]
}

var cases = map[int]int{
	// want: have
	solve(4, []int{1, 2}):    5,
	solve(5, []int{3, 2}):    2,
	solve(4, []int{1, 3, 5}): 3,
}

func main() {

	fmt.Println("testing cases")
	for k, v := range cases {
		if k != v {
			fmt.Println(v)
			fmt.Println("bad case")
			os.Exit(1)
			break
		}
	}
	fmt.Println("success")

}
