package main

import "fmt"

func solve(n, x int) (count int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			res := i * j
			if res == x {
				count++
				break
			}
		}
	}

	return
}

func main() {
	fmt.Println(solve(6, 12))
}
