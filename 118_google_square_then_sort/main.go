package main

import (
	"fmt"
	"sort"
)

func solve(s []int) []int {
	square := []int{}

	for _, v := range s {
		square = append(square, v*v)
	}

	sort.Ints(square)

	return square
}

func main() {
	fmt.Println(solve([]int{-9, -2, 0, 2, 3}))
}
