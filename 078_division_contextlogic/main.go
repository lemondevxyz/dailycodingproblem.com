package main

import "fmt"

func solve(x, y int) int {
	for i := 0; i < x; i++ {
		z := i
		// skip one
		for j := 1; j < y; j++ {
			z += i
		}

		if z == x-1 || z == x+1 || z == x {
			return i
		}
	}

	return -1

	return -1
}

func main() {
	x := map[int]int{
		5:   2,
		10:  2,
		25:  5,
		6:   3,
		100: 10,
	}

	for k, v := range x {
		fmt.Printf("%d / %d = %d\n", k, v, solve(k, v))
	}
}
