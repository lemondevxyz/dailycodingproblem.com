package main

import (
	"fmt"
	"log"
)

func solve(x, y int) int {
	if y == 0 {
		log.Fatalf("error: cannot divide by 0")
	}
	if y == 1 {
		return x
	}
	if x < y {
		return 0
	}
	if x == y {
		return 1
	}

	num := y
	result := 1
	for {
		num += y
		result++
		if num > x {
			result--
			break
		}
	}

	return result
}

func main() {
	xs := []int{48, 43, 44, 26, 16, 30, 34, 19, 43, 4, 49, 47, 39, 31, 42, 15, 46, 40, 14, 31, 35, 30, 14, 43, 40, 6, 46, 31, 39, 24, 1, 31, 42, 45, 46, 41, 30, 11, 23, 21, 26, 49, 35, 30, 40, 41, 31, 46, 35, 17}
	ys := []int{23, 13, 39, 4, 2, 17, 1, 4, 23, 15, 36, 15, 12, 15, 28, 14, 31, 25, 10, 12, 24, 14, 4, 38, 38, 1, 1, 7, 38, 21, 1, 14, 4, 30, 36, 10, 28, 1, 20, 14, 2, 25, 1, 16, 4, 29, 1, 39, 33, 3}

	for index := range xs {
		x, y := xs[index], ys[index]

		if x/y != solve(x, y) {
			fmt.Println("index", index, x, y)
			fmt.Printf("%d / %d = %d\n", x, y, solve(x, y))
			fmt.Println(x / y)
			log.Fatalf("error: division algo is wrong")
		}
	}
}
