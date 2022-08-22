package main

import (
	"fmt"
	"math"
)

var m = map[int]struct{}{
	0: struct{}{},
}

func primeNumbers(n int) {
	for i := 2; i <= n; i++ {
		_, ok := m[i]
		if ok {
			continue
		}

		primal := true

		for j := 2; j < n; j++ {
			if i == j {
				continue
			}

			f := float64(i) / float64(j)
			if math.Mod(f, 1) == 0 {
				primal = false
			}
		}

		if primal {
			m[i] = struct{}{}
		}
	}
}

func solve(x int) (int, int) {
	if x <= 1 {
		return -1, -1
	}

	primeNumbers(x)

	for k1 := range m {
		for k2 := range m {

			if (k1 + k2) == x {
				return k1, k2
			}
		}
	}

	return -1, -1
}

func main() {
	fmt.Println(m)

	for i := 2; i <= 20; i++ {
		x, y := solve(i)
		if x == -1 && y == -1 {
			fmt.Printf("algorithm sucks...\n")
			return
		}

		fmt.Printf("%d + %d = %d\n", x, y, i)
	}
}
