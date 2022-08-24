package main

import (
	"fmt"
)

func exp(x int) int {
	if x <= 1 {
		return 1
	}
	return x * exp(x-1)
}

func solve(s []int) [][]int {
	ret := append([][]int{}, s)
	i, j := 1, 2
	slice := append([]int{}, s...)
	for x := 0; x < exp(len(s))-1; x++ {
		slice[i], slice[j] = slice[j], slice[i]
		ret = append(ret, append([]int{}, slice...))

		i++
		j++

		if j >= len(s) {
			i, j = 0, 1
		}
	}

	return ret
}

func main() {
	fmt.Println(solve([]int{1, 2, 3, 4}))
}
