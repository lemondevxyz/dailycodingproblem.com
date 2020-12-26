package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func solve(multiset []int) bool {

	sums := map[int]bool{}

	sort.Ints(multiset)
	val := 0

	for i := 0; i < len(multiset); i++ {
		val += multiset[i]
		sums[val] = true
	}

	div := float64(val) / 2.00
	if math.Mod(div, 1) == 0 {
		if _, ok := sums[(val / 2)]; ok {
			return true
		}
	}

	return false
}

type Case struct {
	want bool
	have bool
}

var cases = []Case{
	{
		want: true,
		have: solve([]int{15, 5, 20, 10, 35, 15, 10}),
	},
	{
		want: false,
		have: solve([]int{15, 5, 20, 10, 35}),
	},
	{
		want: true,
		have: solve([]int{5, 10, 15, 30}),
	},
}

func main() {
	for k, v := range cases {
		if v.want != v.have {
			fmt.Printf("bad code. index %d\n", k)
			fmt.Println(v.want, v.have)
			os.Exit(1)
		}
	}

	fmt.Println("success")
}
