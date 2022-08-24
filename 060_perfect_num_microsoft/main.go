package main

import "fmt"

func solve(n int) int {
	// idk nibba
	if n > 10 || n <= 0 {
		return -1
	}

	res := n * 10
	add := 10 - n

	return res + add
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i, solve(i))
	}
}
