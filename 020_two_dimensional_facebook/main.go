package main

import (
	"fmt"
	"os"
)

// non-negative lamo
// 0: 2
// 1: 1
// 2: 2
// result 1

// 0: 3
// 1: 0
// 2: 1
// 3: 3
// 4: 0
// 5: 5
// becomes
// 3, 3, 3, 3, 3, 5
// 0, 3, 2, 0, 3, 0
func solve(arr []int) int {
	var count int
	for k := range arr {
		v := arr[k]
		// if it's sandwiched between two
		if k > 0 && k+1 < len(arr) {
			prev := arr[k-1]
			next := arr[k+1]

			compare := -1
			if prev > v {
				compare = prev - v
			} else if next > v {
				compare = next - v
			}

			if compare > 0 {
				count += compare
			}
		}
	}

	return count
}

type Case struct {
	want int
	have int
}

var cases = []Case{
	{
		have: solve([]int{2, 1, 2}),
		want: 1,
	},
	{
		have: solve([]int{3, 0, 1, 3, 0, 5}),
		want: 8,
	},
}

func main() {

	for _, v := range cases {
		if v.have != v.want {
			fmt.Println("bad code")
			os.Exit(1)
		}
	}
	fmt.Println("success")

}
